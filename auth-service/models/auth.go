package models

import (
	"auth-service/constants"
	"auth-service/datatransfers"
	"net/http"
	"strings"

	"gorm.io/gorm"
)

type Auth struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	UserID    string `gorm:"index;unique;type:varchar(12)" json:"userId"`
	Username  string `gorm:"index;unique;type:varchar(255)" validate:"required" json:"username"`
	Password  string `gorm:"type:varchar(255)" validate:"required" json:"-"`
	Token     string `gorm:"-" json:"token"`
	CreatedAt uint   `gorm:"autoCreateTime;<-:create" json:"createdAt"`
	UpdatedAt uint   `gorm:"autoUpdateTime" json:"updatedAt"`
}

func (Auth) TableName() string {
	return "auth"
}

func (a *Auth) setAttr() {
	a.Username = strings.ToLower(strings.TrimSpace(a.Username))
}

func (a *Auth) BeforeCreate(tx *gorm.DB) (err error) {
	// should be handled at max in usecase,
	// so, if this passes through here, we are lacking of validations
	if a.Username == "" || a.Password == "" {
		err = &datatransfers.CustomError{
			Code:    constants.OrmHookDataErrCode,
			Status:  http.StatusInternalServerError,
			Message: "INCOMPLETE_AUTH_DATA",
		}
	}
	a.setAttr()
	return
}

func (a *Auth) BeforeUpdate(tx *gorm.DB) (err error) {
	a.setAttr()
	return
}
