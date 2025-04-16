package models

import (
	"time"

	"gorm.io/gorm"
)

type UserRole string

type User struct {
	ID        uint      `json:"id" gorm:"primarykey"`
	Email     string    `json:"email" gorm:"text;not null"`
	Role      UserRole  `json:"role" gorm:"text;default:attendee"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (u *User) AfterCreate(db *gorm.DB) (err error) {
	u.Role = UserRole("manager")
	return
}
