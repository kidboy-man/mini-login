package models

type User struct {
	ID        string `json:"id"`
	FullName  string `json:"fullName"`
	Email     string `json:"email"`
	CreatedAt uint   `json:"createdAt"`
	UpdatedAt uint   `json:"updatedAt"`
}
