package entity

type User struct {
	ID             string `json:"id"`
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	Password       string `json:"password"`
	Email          string `json:"email"`
	TargetCalories int    `json:"target_calories"`
}
