package database

import (
	"log"
	"time"

	"github.com/spf13/viper"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

// data modules
type Tag struct {
	ID          uint `gorm:"primarykey"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Name        string `gorm:"unique"`
	Description *string

	Group    []TagGroup `gorm:"many2many:taggroup_tags;"`
	Videos   []Video    `gorm:"many2many:video_tags;"`
	Pictures []Picture  `gorm:"many2many:picture_tags;"`
	Audios   []Audio    `gorm:"many2many:audio_tags;"`
	Clips    []Clip     `gorm:"many2many:clip_tags;"`
	Script   []Script   `gorm:"many2many:script_tags;"`
}

type TagGroup struct {
	ID          uint `gorm:"primarykey"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Name        string `gorm:"unique"`
	Description *string
	Tags        []Tag `gorm:"many2many:taggroup_tags;"`
}

type File struct {
	CreatedAt   time.Time
	UpdatedAt   time.Time
	URL         string `gorm:"unique"`
	Size        *int64
	Description *string
	Name        *string
	Hash        *string
}

type Video struct {
	ID uint `gorm:"primarykey"`
	File
	Duriation uint
	Clips     []Clip `gorm:"foreignKey:VideoID"`
	Tags      []Tag  `gorm:"many2many:video_tags;"`
}

// child of video
type Clip struct {
	ID          uint `gorm:"primarykey"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Name        *string
	Description *string
	Start       uint
	End         uint
	Tags        []Tag `gorm:"many2many:clip_tags;"`

	VideoID uint `gorm:"constraint:OnDelete:CASCADE;"`
}

type Picture struct {
	ID uint `gorm:"primarykey"`
	File
	Tags []Tag `gorm:"many2many:picture_tags;"`
}

type Audio struct {
	ID uint `gorm:"primarykey"`
	File
	Tags []Tag `gorm:"many2many:audio_tags;"`
}

type Script struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Tags      []Tag  `gorm:"many2many:script_tags;"`
	Text      string `gorm:"type:text"`
	Name      string `gorm:"unique"`
}

func InitDB() *gorm.DB {
	ldb, err := gorm.Open(sqlite.Open(viper.GetString("database")), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	err = ldb.AutoMigrate(&Tag{}, &TagGroup{}, &Video{}, &Clip{}, &Picture{}, &Audio{}, &Script{})
	if err != nil {
		log.Fatalln(err)
	}
	DB = ldb
	return ldb
}
