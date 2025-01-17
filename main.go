package main

//go:generate swag init

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"time"

	"github.com/spf13/viper"

	"slidegoose/config"
	"slidegoose/internal/database"
	"slidegoose/internal/router"
)

// @title           Slide Goose
// @version         alpha
// @license.name 	GPL-3.0-or-later
// @license.url		https://www.gnu.org/licenses/gpl-3.0.html
// @BasePath  /api
func main() {
	// chdir to right place
	execPath, err := os.Executable()
	if err != nil {
		log.Fatalln(err)
	}
	execPath = filepath.Dir(execPath)
	os.Chdir(execPath)

	// read config
	config.ReadConfig()

	database.InitDB()

	// get router
	router := router.GetRouter()

	// log the address
	log.Printf("open on address %s:%s", viper.GetString("host"), viper.GetString("port"))

	// define server
	srv := &http.Server{
		Addr:    viper.GetString("host") + ":" + viper.GetString("port"),
		Handler: router,
	}

	// run server
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// graceful exit server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")

}
