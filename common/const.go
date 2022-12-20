package common

type DBType int

const (
	DBTypeRestaurant DBType = 1
	DBTypeUser       DBType = 2
)

const (
	CurrentUser = "user"
)

type Requester interface {
	GetUserId() int
	GetEmail() string
	GetRole() string
}
