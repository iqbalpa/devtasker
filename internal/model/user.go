package model

import "time"

type User struct {
	ID           string    `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Name         string    `json:"name"`
	Username     string    `json:"username" gorm:"unique"`
	PasswordHash string    `json:"password"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`

	Tasks []Task `gorm:"foreignKey:UserUsername;references:Username"`
}
