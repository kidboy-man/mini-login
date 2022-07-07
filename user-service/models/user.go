package models

import (
	"strings"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Username  string    `gorm:"index;unique;type:varchar(255)" validate:"required" json:"username"`
	Password  string    `gorm:"type:varchar(255)" validate:"required" json:"password"`
	Email     string    `gorm:"index;unique;type:varchar(255)" validate:"required,email" json:"email"`
	Token     string    `gorm:"-" json:"token"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (User) TableName() string {
	return "user"
}

func (u *User) setAttr() {
	u.Username = strings.ToLower(strings.TrimSpace(u.Username))
	u.Email = strings.ToLower(strings.TrimSpace(u.Email))
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.setAttr()
	return
}

func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
	u.setAttr()
	return
}
