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

type tagable interface {
	SetTag(Tag) error
}

func (s *Video) SetTag(tag Tag) error {
	s.Tags = append(s.Tags, tag)
	if err := db.Save(&s).Error; err != nil {
		return err
	}
	return nil
}

func (s *Clip) SetTag(tag Tag) error {
	s.Tags = append(s.Tags, tag)
	if err := db.Save(&s).Error; err != nil {
		return err
	}
	return nil
}
func (s *Picture) SetTag(tag Tag) error {
	s.Tags = append(s.Tags, tag)
	if err := db.Save(&s).Error; err != nil {
		return err
	}
	return nil
}
func (s *Audio) SetTag(tag Tag) error {
	s.Tags = append(s.Tags, tag)
	if err := db.Save(&s).Error; err != nil {
		return err
	}
	return nil
}
func (s *Script) SetTag(tag Tag) error {
	s.Tags = append(s.Tags, tag)
	if err := db.Save(&s).Error; err != nil {
		return err
	}
	return nil
}

func SetTag(s tagable, tag_id uint) error {
	var tag Tag
	if err := db.First(&tag, tag_id).Error; err != nil {
		return err
	}
	if err := s.SetTag(tag); err != nil {
		return err
	}
	return nil
}

func SetVideoTag(id uint, tag_id uint) error {
	var s Video
	if err := db.First(&s, id).Error; err != nil {
		return err
	}

	if err := SetTag(&s, tag_id); err != nil {
		return err
	}
	return nil
}

func SetClipTag(id uint, tag_id uint) error {
	var s Clip
	if err := db.First(&s, id).Error; err != nil {
		return err
	}

	if err := SetTag(&s, tag_id); err != nil {
		return err
	}
	return nil
}

func SetPictureTag(id uint, tag_id uint) error {
	var s Picture
	if err := db.First(&s, id).Error; err != nil {
		return err
	}

	if err := SetTag(&s, tag_id); err != nil {
		return err
	}
	return nil
}

func SetAudioTag(id uint, tag_id uint) error {
	var s Audio
	if err := db.First(&s, id).Error; err != nil {
		return err
	}

	if err := SetTag(&s, tag_id); err != nil {
		return err
	}
	return nil
}

func SetScriptTag(id uint, tag_id uint) error {
	var s Script
	if err := db.First(&s, id).Error; err != nil {
		return err
	}

	if err := SetTag(&s, tag_id); err != nil {
		return err
	}
	return nil
}
