package usermodel

import (
	"errors"
	"golang_01/common"
	"strings"
)

const EntityName = "User"

type User struct {
	common.SQLModel
	Email     string          `json:"email" gorm:"column:email;"`
	Password  string          `json:"-" gorm:"column:password;"`
	Salt      string          `json:"-" gorm:"column:salt;"`
	LastName  string          `json:"last_name" gorm:"column:last_name;"`
	FirstName string          `json:"first_name" gorm:"column:first_name;"`
	Phone     string          `json:"phone" gorm:"column:phone;"`
	Role      common.UserRole `json:"role" gorm:"column:role;"`
	Avatar    *common.Image   `json:"avatar,omitempty" gorm:"column:avatar;"`
}

func (data *User) GetUserId() int {
	return data.Id
}

func (data *User) GetStatus() int {
	return data.Status
}

func (data *User) GetEmail() string {
	return data.Email
}

func (data *User) GetRole() string {
	return data.Role.String()
}

func (User) TableName() string {
	return "users"
}

type UserCreate struct {
	common.SQLModel
	Email     string        `json:"email" gorm:"column:email;"`
	Password  string        `json:"password" gorm:"column:password;"`
	Salt      string        `json:"salt" gorm:"column:salt;"`
	LastName  string        `json:"lastName" gorm:"column:last_name;"`
	FirstName string        `json:"firstName" gorm:"column:first_name;"`
	Phone     int           `json:"phone" gorm:"column:phone;"`
	Role      string        `json:"role" gorm:"column:role;"`
	Avatar    *common.Image `json:"avatar,omitempty" gorm:"column:avatar;"`
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

type UserUpdate struct {
	LastName  string        `json:"lastName" gorm:"column:last_name;"`
	FirstName string        `json:"firstName" gorm:"column:first_name;"`
	Phone     int           `json:"phone" gorm:"column:phone;"`
	Avatar    *common.Image `json:"avatar,omitempty" gorm:"column:avatar;"`
}

func (UserUpdate) TableName() string {
	return User{}.TableName()
}

type UserUpdatePassWord struct {
	OldPassword   string `json:"oldPassword" gorm:"-"`
	MatchPassWord string `json:"matchPassWord" gorm:"-"`
	NewPassword   string `json:"newPassword" gorm:"column:password;"`
	Salt          string `json:"salt" gorm:"column:salt;"`
}

func (UserUpdatePassWord) TableName() string {
	return User{}.TableName()
}

type UserUpdateStatus struct {
	Status int `json:"status" gorm:"status;"`
	UserId int `json:"userId" gorm:"user_id;"`
}

func (UserUpdateStatus) TableName() string {
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
	ErrPassWrong = common.NewCustomError(
		errors.New("you password is wrong"),
		"you password is wrong",
		"ErrPassWrong",
	)
	ErrMathPassWrong = common.NewCustomError(
		errors.New("you password not match"),
		"you password not match",
		"ErrPassNotMatch",
	)
	ErrUpdateStatus = common.NewCustomError(
		errors.New("you can not update your status"),
		"you password not match",
		"ErrPassNotMatch",
	)
)

func (data *User) Mask(isAdminOrOwner bool) {
	data.GenUID(common.DBTypeUser)
}

func (res *UserUpdatePassWord) Validate() error {
	res.OldPassword = strings.TrimSpace(res.OldPassword)
	res.NewPassword = strings.TrimSpace(res.NewPassword)
	res.MatchPassWord = strings.TrimSpace(res.MatchPassWord)

	if len(res.OldPassword) == 0 {
		return errors.New("password can not be blank")
	}

	if len(res.NewPassword) == 0 {
		return errors.New("new password can not be blank")
	}

	if len(res.MatchPassWord) == 0 {
		return errors.New("match password can not be blank")
	}

	if res.NewPassword != res.MatchPassWord {
		return errors.New("password not match")
	}

	return nil
}
