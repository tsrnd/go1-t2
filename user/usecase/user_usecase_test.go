package usecase_test

import (
	userRepo "goweb2/user/repository"
	userUse  "goweb2/user/usecase"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestGetByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	rows := sqlmock.NewRows([]string{"id", "name", "email", "phone", "created_at", "updated_at"}).
		AddRow(10, "name", "email@gmail.com", "0906777888", time.Now(), time.Now())

	query := "select id, name, email, phone, created_at, updated_at from users where id = ?"

	mock.ExpectQuery(query).WithArgs(10).WillReturnRows(rows)
	a := userRepo.NewUserRepository(db)
	b := userUse.NewUserUsecase(a)
	num := int64(10)
	anArticle, err := b.GetByID(num)
	assert.NoError(t, err)
	assert.NotNil(t, anArticle)
}
