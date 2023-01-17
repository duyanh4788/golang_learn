package restaurantmodel

import (
	"errors"
	"golang_01/common"
	"strings"
)

const EntityName = "restaurants"

type Restaurants struct {
	common.SQLModel `json:",inline"`
	OwnerId         int            `json:"owner_id" gorm:"column:owner_id;"`
	Name            string         `json:"name" gorm:"column:name;"`
	Addr            string         `json:"address" gorm:"column:addr;"`
	Logo            *common.Image  `json:"logo" gorm:"column:logo;"`
	Cover           *common.Images `json:"cover" gorm:"column:cover;"`
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
	Id      int            `json:"id" gorm:"column:id;"`
	Name    string         `json:"name" gorm:"column:name;"`
	OwnerId int            `json:"-" gorm:"column:owner_id"`
	Addr    string         `json:"address" gorm:"column:addr;"`
	Logo    *common.Image  `json:"logo" gorm:"column:logo;"`
	Cover   *common.Images `json:"cover" gorm:"column:cover;"`
}

func (RestaurantCreate) TableName() string {
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
