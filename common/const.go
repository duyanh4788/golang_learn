package common

const (
	DBTypeRestaurant = 1
	DBTypeUser       = 2
	DBTypeCategory   = 3
)

const (
	CurrentUser = "user"
)

const (
	TimeLayout = "2006-01-02T15:04:05.999999"
)

type Requester interface {
	GetUserId() int
	GetEmail() string
	GetRole() string
}
