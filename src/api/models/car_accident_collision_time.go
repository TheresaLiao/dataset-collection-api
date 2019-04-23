package models

type CarAccidentCollisionTime struct {
	ID            uint   `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	CarAccidentID uint   `json:"car_accident_id"`
	CollisionTime string `json:"collision_time"`
}
