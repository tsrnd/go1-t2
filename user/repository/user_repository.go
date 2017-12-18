package repository

import (
	"database/sql"

	model "goweb2/user"
)

// UserRepository interface
type UserRepository interface {
	GetByID(id int64) (*model.User, error)
}

// userRepository struct
type userRepository struct {
	DB *sql.DB
}

// GetByID func
func (m *userRepository) GetByID(id int64) (*model.User, error) {
	const query = `SELECT id, name, email, phone FROM users WHERE id = $1`
	var user model.User
	err := m.DB.QueryRow(query, id).Scan(&user.ID,  &user.Name, &user.Email, &user.Phone)
	return &user, err
}

// NewUserRepository func
func NewUserRepository(DB *sql.DB) UserRepository {
	return &userRepository{DB}
}
