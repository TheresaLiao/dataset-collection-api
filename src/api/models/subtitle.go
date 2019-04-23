package models

type Subtitle struct {
	ID               uint   `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	Title            string `json:"title"`
	VideoLanguage    string `json:"video_language"`
	SubtitleLanguage string `json:"subtitle_language"`
	URL              string `json:"url"`
	CopyRight        string `json:"copy_right"`
	VideoID          string `json:"video_id"`
	YoutubeID        string `json:"youtube_id"`
	Embedded         bool   `json:"embedded"`
	PlugIn           bool   `json:"plug_in"`
	VideoLength      string `json:"video_length"`
}
