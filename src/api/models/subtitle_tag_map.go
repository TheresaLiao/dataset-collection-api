package models

type SubtitleTagMap struct {
	ID            uint `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	SubtitleID    uint `json:"subtitle_id"`
	SubtitleTagID uint `json:"subtitle_tag_id"`
}
