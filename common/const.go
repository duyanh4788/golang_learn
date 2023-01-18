package common

const (
	DBTypeRestaurant = 1
	DBTypeUser       = 2
	DBTypeCategory   = 3
)

const (
	CurrentUser = "user"
)

type Requester interface {
	GetUserId() int
	GetEmail() string
	GetRole() string
}
