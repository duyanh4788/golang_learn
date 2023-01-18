package restaurantlikemodel

import "time"

type RestaurantLike struct {
	RestaurantId int        `json:"restaurant_id" gorm:"column:restaurant_id;"`
	UserId       int        `json:"user_id" gorm:"column:user_id;"`
	CreateAt     *time.Time `json:"create_at" gorm:"column:create_at;"`
}

func (RestaurantLike) TableName() string {
	return "restaurant_likes"
}
