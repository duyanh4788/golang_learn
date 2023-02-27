package restaurantmodel

import (
	"errors"
	"golang_01/common"
	"strings"
)

const EntityName = "restaurants"

type Restaurants struct {
	common.SQLModel `json:",inline"`
	UserId          int                `json:"owner_id" gorm:"column:owner_id;"`
	Name            string             `json:"name" gorm:"column:name;"`
	Addr            string             `json:"address" gorm:"column:addr;"`
	Logo            *common.Image      `json:"logo" gorm:"column:logo;"`
	Cover           *common.Images     `json:"cover" gorm:"column:cover;"`
	User            *common.SimpleUser `json:"user" gorm:"preload:false;"`
	LikeCount       int                `json:"like_count" gorm:"column:like_count;"` // computed field
	Liked           bool               `json:"liked" gorm:"-";`
}

func (Restaurants) TableName() string {
	return "restaurants"
}

type RestaurantUpdate struct {
	Name  *string        `json:"name" gorm:"column:name;"`
	Addr  *string        `json:"address" gorm:"column:addr;"`
	Logo  *common.Image  `json:"logo" gorm:"column:logo;"`
	Cover *common.Images `json:"cover" gorm:"column:cover;"`
}

func (RestaurantUpdate) TableName() string {
	return Restaurants{}.TableName()
}

type RestaurantCreate struct {
	common.SQLModel `json:",inline"`
	UserId          int            `json:"-" gorm:"column:owner_id"`
	Name            string         `json:"name" gorm:"column:name;"`
	Addr            string         `json:"address" gorm:"column:addr;"`
	Logo            *common.Image  `json:"logo" gorm:"column:logo;"`
	Cover           *common.Images `json:"cover" gorm:"column:cover;"`
}

func (RestaurantCreate) TableName() string {
	return Restaurants{}.TableName()
}

type UpdateStatusRestaurant struct {
	Status       int `json:"status" gorm:"status;"`
	RestaurantId int `json:"restaurantId" gorm:"restaurant_id;"`
}

func (UpdateStatusRestaurant) TableName() string {
	return Restaurants{}.TableName()
}

func (res *RestaurantCreate) Validate() error {
	res.Name = strings.TrimSpace(res.Name)
	res.Addr = strings.TrimSpace(res.Addr)

	if len(res.Name) == 0 {
		return errors.New("restaurant name can not be blank")
	}

	if len(res.Addr) == 0 {
		return errors.New("address can not be blank")
	}

	return nil
}

func (data *Restaurants) Mask(isAdminOrOwner bool) {
	data.GenUID(common.DBTypeRestaurant)

	if u := data.User; u != nil {
		u.Mask(isAdminOrOwner)
	}
}

func (data *RestaurantCreate) MaskCreate(isAdminOrOwner bool) {
	data.GenUID(common.DBTypeRestaurant)
}
