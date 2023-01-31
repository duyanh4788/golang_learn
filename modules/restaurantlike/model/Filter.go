package restaurantlikemodel

type Filter struct {
	RestaurantId int    `json:"-" form:"restaurant_id"`
	UserId       string `json:"-" from:"user_id";`
}
