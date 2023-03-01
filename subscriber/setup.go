package subscriber

import (
	"context"
	"golang_01/component"
)

func Setup(appContext component.AppContext) {
	IncreaseUserLikedRestaurant(appContext, context.Background())
	DecreaseUserUnlikedRestaurant(appContext, context.Background())
}
