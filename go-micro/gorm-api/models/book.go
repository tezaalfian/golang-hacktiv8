package models

import "time"

type Book struct {
	ID        uint   `gorm:"primary_key"`
	NameBook  string `gorm:"type:varchar(191);not null;unique" json:"name_book"`
	Author    string `gorm:"type:varchar(191);not null" json:"author"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
