package usermodel

import (
	"errors"
	"golang_01/common"
)

const EntityName = "User"

type User struct {
	common.SQLModel
	Email    string          `json:"email" gorm:"column:email;"`
	Password string          `json:"-" gorm:"column:password;"`
	Salt     string          `json:"-" gorm:"column:salt;"`
	LastName string          `json:"last_name" gorm:"column:last_name;"`
	FistName string          `json:"fist_name" gorm:"column:fist_name;"`
	Phone    string          `json:"phone" gorm:"column:phone;"`
	Role     common.UserRole `json:"role" gorm:"column:role;"`
	Avatar   *common.Image   `json:"avatar,omitempty" gorm:"column:avatar;"`
}

func (u *User) GetUserId() int {
	return u.Id
}

func (u *User) GetEmail() string {
	return u.Email
}

func (u *User) GetRole() string {
	return u.Role.String()
}

func (User) TableName() string {
	return "users"
}

type UserCreate struct {
	common.SQLModel
	Email    string        `json:"email" gorm:"column:email;"`
	Password string        `json:"password" gorm:"column:password;"`
	Salt     string        `json:"salt" gorm:"column:salt;"`
	LastName string        `json:"lastName" gorm:"column:last_name;"`
	FistName string        `json:"fistName" gorm:"column:fist_name;"`
	Phone    int           `json:"phone" gorm:"column:phone;"`
	Role     string        `json:"role" gorm:"column:role;"`
	Avatar   *common.Image `json:"avatar,omitempty" gorm:"column:avatar;"`
}

func (UserCreate) TableName() string {
	return User{}.TableName()
}

type UserLogin struct {
	Email    string `json:"email" gorm:"column:email;"`
	Password string `json:"password" gorm:"column:password;"`
}

func (UserLogin) TableName() string {
	return User{}.TableName()
}

var (
	ErrEmailOrPassWordInvalid = common.NewCustomError(
		errors.New("email or password invalid"),
		"email or password invalid",
		"ErrEmailOrPassWordInvalid",
	)
	ErrEmailExited = common.NewCustomError(
		errors.New("email has already exits"),
		"email has already exits",
		"ErrEmailExited",
	)
)
