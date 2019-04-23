package models

type CarAccidentTagMap struct {
	ID               uint            `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	CarAccidentID    uint            `json:"car_accident_id"`
	CarAccidentTagID uint            `json:"car_accident_tag_id"`
	CarAccidentTag   *CarAccidentTag `json:"car_accident_tag"` // This line is infered from column name "car_accident_tag_id".

}
