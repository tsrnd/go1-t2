package user

import (
	"time"
)

// User struct
type User struct {
	ID         int64     `json:"id"`
	Name       string    `json:"name"`
	Email      string    `json:"email"`
	Phone      string    `json:"phone"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
}

// PrivateUserDetails struct
type PrivateUserDetails struct {
	ID       int64
	Password string
	Token    string
}
