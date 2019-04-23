package models

type CarAccidentTag struct {
	ID                uint                 `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	TagName           string               `json:"tag_name"`
	CarAccidentTagMap []*CarAccidentTagMap `json:"car_accident_tag_map"` // This line is infered from other tables.
}
