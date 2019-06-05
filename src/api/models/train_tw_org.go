package models

type TrainTwOrg struct {
	Caraccidentid string `json:"CarAccidentID"`
	Title         string `json:"title"`
	Copyright     string `json:"CopyRight"`
	URL           string `json:"URL"`
	Keyword       string `json:"KeyWord"`
	VideoLength   string `json:"video_length"`
	CollisionTime string `json:"collision_time"`
	TagName       string `json:"tag_name"`
	CarType       string `json:"car_type"`
}
