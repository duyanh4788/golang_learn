package restaurantlikemodel

import (
	"golang_01/common"
	"time"
)

const EntityName = "restaurant_likes"

type RestaurantLike struct {
	RestaurantId int                `json:"restaurant_id" gorm:"column:restaurant_id;"`
	UserId       int                `json:"user_id" gorm:"column:user_id;"`
	CreatedAt    *time.Time         `json:"created_at" gorm:"column:created_at;"`
	User         *common.SimpleUser `json:"user" gorm:"preload:false";`
}

func (RestaurantLike) TableName() string {
	return "restaurant_likes"
}
