package entity

import "time"

type User struct {
	ID             string    `json:"id"`
	FirstName      string    `json:"first_name"`
	LastName       string    `json:"last_name"`
	Password       string    `json:"password"`
	Email          string    `json:"email"`
	TargetCalories int       `json:"target_calories"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
