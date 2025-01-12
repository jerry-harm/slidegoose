package database

import (
	"database/sql"
	"log"
	"time"

	"github.com/spf13/viper"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

// data modules
type Tag struct {
	ID          uint `gorm:"primarykey"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Name        string `gorm:"primarykey"`
	Description sql.NullString
	Group       []*TagGroup `gorm:"many2many:taggroup_tags;"`

	Videos   []*Tag `gorm:"many2many:video_tags;"`
	Pictures []*Tag `gorm:"many2many:picture_tags;"`
	Audios   []*Tag `gorm:"many2many:audio_tags;"`
	Clips    []*Tag `gorm:"many2many:clip_tags;"`
}

type TagGroup struct {
	ID          uint `gorm:"primarykey"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Name        string `gorm:"primarykey"`
	Description sql.NullString
	Tags        []*Tag `gorm:"many2many:taggroup_tags;"`
}

type Video struct {
	ID          uint `gorm:"primarykey"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Path        string `gorm:"unique"`
	Type        string
	Size        int64
	ModTime     time.Time
	Description sql.NullString
	Tags        []*Tag `gorm:"many2many:video_tags;"`
	Duriation   uint
	Clips       []Clip `gorm:"foreignKey:VideoID"`
}

// child of video
type Clip struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Start     uint
	End       uint
	Tags      []*Tag `gorm:"many2many:clip_tags;"`

	VideoID uint `gorm:"constraint:OnDelete:CASCADE;"`
}

type Picture struct {
	ID          uint `gorm:"primarykey"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Path        string `gorm:"unique"`
	Type        string
	Size        int64
	ModTime     time.Time
	Description sql.NullString
	Tags        []*Tag `gorm:"many2many:picture_tags;"`
}

type Audio struct {
	ID          uint `gorm:"primarykey"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Path        string `gorm:"unique"`
	Type        string
	Size        int64
	ModTime     time.Time
	Description sql.NullString
	Tags        []*Tag `gorm:"many2many:audio_tags;"`
}

type Script struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Text      string `gorm:"type:text"`
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
	db = ldb
	return ldb
}
