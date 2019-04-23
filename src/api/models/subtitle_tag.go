package models

type SubtitleTag struct {
	ID      uint   `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	TagName string `json:"tag_name"`
}
