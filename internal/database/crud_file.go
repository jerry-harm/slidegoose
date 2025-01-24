package database

import (
	"io/fs"
	"log"
	"net/url"
	"path/filepath"
	"slices"
)

// TODO only in local dir? or allow to pack up all the things...

var videoTypes = []string{".mp4"}

var pictureTypes = []string{".jpg", ".jepg", ".png"}

var audioTypes = []string{".mp3"}

func AddLocalFile(path string, d fs.DirEntry) bool {
	info, err := d.Info()
	if err != nil {
		return false
	}
	var url url.URL
	url.Path = path
	url.Scheme = "file"
	ext := filepath.Ext(path)
	size := info.Size()
	if slices.Contains(videoTypes, ext) {
		value := Audio{}
		value.URL = url.String()
		value.Size = &size
		if err := DB.Create(&value).Error; err != nil {
			return false
		}
		return true
	} else if slices.Contains(pictureTypes, ext) {
		value := Picture{}
		value.URL = url.String()
		value.Size = &size
		if err := DB.Create(&value).Error; err != nil {
			return false
		}
		return true
	} else if slices.Contains(audioTypes, ext) {
		value := Audio{}
		value.URL = url.String()
		value.Size = &size
		if err := DB.Create(&value).Error; err != nil {
			return false
		}
		return true
	}

	return false
}

func AddDir(path string) int {
	var files_count int = 0
	filepath.WalkDir(path, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			log.Println(err)
			return filepath.SkipDir
		}

		if !d.IsDir() {
			if AddLocalFile(path, d) {
				files_count++
			}

		}

		return nil
	})

	return files_count
}
