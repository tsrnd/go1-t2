package repository

import (
	"database/sql"

	model "goweb2/user"
)

// userRepository struct
type userRepository struct {
	DB *sql.DB
}

// NewUserRepository func
func NewUserRepository(DB *sql.DB) UserRepository {
	return &userRepository{DB}
}

// GetByID func
func (m *userRepository) GetByID(id int64) (*model.User, error) {
	const query = `
    select
      id,
      name,
      email,
      phone,
	  created_at,
	  updated_at
    from
      users
    where
      id = $1
  `
	var user model.User
	err := m.DB.QueryRow(query, id).Scan(&user.ID, &user.Name, &user.Email, &user.Phone, &user.Created_at, &user.Updated_at)
	return &user, err
}
