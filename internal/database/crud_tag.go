package database

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

func SetVideoTag(id uint, tag_id uint) error {
	var tag Tag

	if err := db.First(&tag, tag_id).Error; err != nil {
		return err
	}

	var s Video
	if err := db.First(&s, id).Error; err != nil {
		return err
	}

	s.Tags = append(s.Tags, &tag)
	if err := db.Save(&s).Error; err != nil {
		return err
	}
	return nil
}

func SetClipTag(id uint, tag_id uint) error {
	var tag Tag

	if err := db.First(&tag, tag_id).Error; err != nil {
		return err
	}

	var s Clip
	if err := db.First(&s, id).Error; err != nil {
		return err
	}

	s.Tags = append(s.Tags, &tag)
	if err := db.Save(&s).Error; err != nil {
		return err
	}
	return nil
}

func SetPictureTag(id uint, tag_id uint) error {
	var tag Tag

	if err := db.First(&tag, tag_id).Error; err != nil {
		return err
	}

	var s Picture
	if err := db.First(&s, id).Error; err != nil {
		return err
	}

	s.Tags = append(s.Tags, &tag)
	if err := db.Save(&s).Error; err != nil {
		return err
	}
	return nil
}

func SetAudioTag(id uint, tag_id uint) error {
	var tag Tag

	if err := db.First(&tag, tag_id).Error; err != nil {
		return err
	}

	var s Audio
	if err := db.First(&s, id).Error; err != nil {
		return err
	}

	s.Tags = append(s.Tags, &tag)
	if err := db.Save(&s).Error; err != nil {
		return err
	}
	return nil
}
