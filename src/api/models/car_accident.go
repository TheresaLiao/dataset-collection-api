package models

type CarAccident struct {
	ID          uint   `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	Title       string `json:"title"`
	URL         string `json:"url"`
	CopyRight   string `json:"copy_right"`
	VideoLength string `json:"video_length"`
	CarType     string `json:"car_type"`
	YoutubeID   string `json:"youtube_id"`
}
