package database

import (
	"io/fs"
	"log"
	"path/filepath"
	"slices"

	"gorm.io/gorm"
)

var videoTypes = []string{".mp4"}

var pictureTypes = []string{".jpg", ".jepg", ".png"}

var audioTypes = []string{".mp3"}

func AddVideo(db *gorm.DB) {}

func AddPicture(db *gorm.DB) {}

func AddAudio(db *gorm.DB) {}

func AddFile(db *gorm.DB, path string, d fs.DirEntry) error {
	info, err := d.Info()
	if err != nil {
		log.Panicln(path, err, "not add")
	} else {
		ext := filepath.Ext(path)
		info.ModTime()
		if slices.Contains(videoTypes, ext) {
			AddVideo(db)
		} else if slices.Contains(pictureTypes, ext) {
			AddPicture(db)
		} else if slices.Contains(audioTypes, ext) {
			AddAudio(db)
		}
	}
	return nil
}

func AddDir(db *gorm.DB, path string) {
	filepath.WalkDir(path, func(path string, d fs.DirEntry, err error) error {
		if !d.IsDir() {
			err := AddFile(db, path, d)
			if err != nil {

			}
		}
		return nil
	})
}
