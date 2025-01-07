package internal

import (
	"database/sql"
	"log"
	"time"

	"github.com/spf13/viper"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Tag struct {
	ID          uint
	Name        string
	Description sql.NullString
}

type File struct {
	ID          uint
	Path        string
	Type        string
	Description sql.NullString
	Tags        []Tag
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func InitDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open(viper.GetString("database")), &gorm.Config{})
	if err != nil {
		log.Println(err)
		// create db
	}
	return db
}
