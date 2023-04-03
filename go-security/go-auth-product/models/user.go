package models

import (
	"github.com/asaskevich/govalidator"
	"go-auth-product/helpers"
	"gorm.io/gorm"
)

type Role string

type User struct {
	GormModel
	FullName string    `gorm:"not null" json:"full_name" form:"full_name" valid:"required"`
	Email    string    `gorm:"not null;uniqueIndex" json:"email" form:"email" valid:"required,email"`
	Password string    `gorm:"not null" json:"password" form:"password" valid:"required,minstringlength(6)"`
	Role     string    `gorm:"not null" json:"role"`
	Products []Product `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"products"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)
	if errCreate != nil {
		err = errCreate
		return
	}
	u.Password = helpers.HashPass(u.Password)
	u.Role = "user"
	err = nil
	return
}
