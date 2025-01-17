package database

import (
	"io/fs"
	"log"
	"path/filepath"
	"slices"
)

// TODO only in local dir? or allow to pack up all the things...

var videoTypes = []string{".mp4"}

var pictureTypes = []string{".jpg", ".jepg", ".png"}

var audioTypes = []string{".mp3"}

func AddVideo(path string, info fs.FileInfo) {
	value := Video{
		Path:    path,
		Size:    info.Size(),
		Type:    filepath.Ext(path),
		ModTime: info.ModTime(),
	}
	db.Create(&value)
}

func AddPicture(path string, info fs.FileInfo) {
	value := Picture{
		Path:    path,
		Size:    info.Size(),
		Type:    filepath.Ext(path),
		ModTime: info.ModTime(),
	}
	db.Create(&value)
}

func AddAudio(path string, info fs.FileInfo) {
	value := Audio{
		Path:    path,
		Size:    info.Size(),
		Type:    filepath.Ext(path),
		ModTime: info.ModTime(),
	}
	db.Create(&value)
}

func AddFile(path string, d fs.DirEntry) bool {
	info, err := d.Info()
	if err != nil {
		log.Println(path, err)
	} else {
		ext := filepath.Ext(path)
		info.ModTime()
		if slices.Contains(videoTypes, ext) {
			AddVideo(path, info)
			return true
		} else if slices.Contains(pictureTypes, ext) {
			AddPicture(path, info)
			return true
		} else if slices.Contains(audioTypes, ext) {
			AddAudio(path, info)
			return true
		}

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
			if AddFile(path, d) {
				files_count++
			}

		}

		return nil
	})

	return files_count
}
