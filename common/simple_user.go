package common

type SimpleUser struct {
	SQLModel  `json:",inline"`
	LastName  string `json:"last_name" gorm:"column:last_name;"`
	FirstName string `json:"first_name" gorm:"column:first_name;"`
	Role      string `json:"role" gorm:"column:role;"`
}

func (SimpleUser) TableName() string {
	return "users"
}

func (data *SimpleUser) Mask(isAdminOrOwner bool) {
	data.GenUID(DBTypeUser)
}
