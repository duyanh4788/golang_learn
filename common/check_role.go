package common

import (
	"database/sql/driver"
	"errors"
	"fmt"
)

type UserRole int

const (
	RoleUser UserRole = 1 << iota
	RoleAdmin
	RoleShipper
	RoleMod
)

func (role UserRole) String() string {
	switch role {
	case RoleAdmin:
		return "admin"
	case RoleShipper:
		return "shipper"
	case RoleMod:
		return "mod"
	default:
		return "user"
	}
}

func (role *UserRole) Scan(value interface{}) error {
	bytes, ok := value.([]byte)

	if !ok {
		errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}

	var roleType UserRole

	result := string(bytes)

	switch result {
	case "admin":
		roleType = RoleUser
	case "shipper":
		roleType = RoleShipper
	case "mod":
		roleType = RoleMod
	default:
		roleType = RoleUser
	}

	*role = roleType

	return nil
}

func (role *UserRole) Value() (driver.Value, error) {
	if role == nil {
		return nil, nil
	}
	return role.String(), nil
}
func (role *UserRole) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%s\"", role.String())), nil
}
