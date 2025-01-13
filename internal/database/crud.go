package database

import (
	"io/fs"
	"log"
	"path/filepath"
	"slices"
)

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

func AddFile(path string, d fs.DirEntry) {
	info, err := d.Info()
	if err != nil {
		log.Println(path, err)
	} else {
		ext := filepath.Ext(path)
		info.ModTime()
		if slices.Contains(videoTypes, ext) {
			AddVideo(path, info)
		} else if slices.Contains(pictureTypes, ext) {
			AddPicture(path, info)
		} else if slices.Contains(audioTypes, ext) {
			AddAudio(path, info)
		}
	}
}

func AddDir(path string) int {
	var files_count int = 0
	filepath.WalkDir(path, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			log.Println(err)
			return filepath.SkipDir
		}

		if !d.IsDir() {
			AddFile(path, d)
			files_count++
		}

		return nil
	})

	return files_count
}

func AddTag(name string, description string) {
	value := Tag{
		Name: name,
	}
	if len(description) != 0 {
		value.Description.String = description
		value.Description.Valid = true
	}

	db.Create(&value)
}

func SetVideoTag(video_id uint, tag_id uint) error {
	var video Video
	if err := db.First(&video, video_id).Error; err != nil {
		return err
	}

	var tag Tag

	if err := db.First(&tag, tag_id).Error; err != nil {
		return err
	}
	video.Tags = append(video.Tags, &tag)
	if err := db.Save(&video).Error; err != nil {
		return err
	}
	return nil
}
