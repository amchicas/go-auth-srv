package mysql

import (
	"context"
	"database/sql"
	"errors"

	"github.com/amchicas/go-auth-srv/internal/domain"
)

const (
	QUERY_GET_USERS            = "SELECT * FROM auth"
	QUERY_GET_USER             = "SELECT * FROM auth WHERE id = ? "
	QUERY_GET_USER_EMAIL       = "SELECT id,username,email,password,role FROM auth WHERE email  like  ? limit 1 "
	QUERY_CREATE_USER          = "INSERT INTO auth (username, email,password,role, created, modified) VALUES (?, ?, ?, ?,?,?)"
	QUERY_UPDATE_USER          = "UPDATE auth SET username = ?,role = ? , modified = ? WHERE id = ?"
	QUERY_UPDATE_USER_PASSWORD = "UPDATE auth SET password = ? , modified = ? WHERE id = ?"
	QUERY_DELETE_USER          = "DELETE FROM auth WHERE id = ?"
)

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) domain.Repository {
	return &repository{
		db: db,
	}
}
func (r *repository) CreateUser(ctx context.Context, user *domain.Auth) error {
	stmt, err := r.db.PrepareContext(ctx, QUERY_CREATE_USER)
	if err != nil {

		return err

	}
	defer stmt.Close()
	_, err = stmt.ExecContext(ctx, user.Username, user.Email, user.Password, user.Role, user.Modified, user.Created)
	if err != nil {
		return err

	}
	return nil

}

func (r *repository) GetByEmail(ctx context.Context, email string) (*domain.Auth, error) {
	auth := &domain.Auth{}
	stmt, err := r.db.PrepareContext(ctx, QUERY_GET_USER_EMAIL)
	if err != nil {
		return nil, err

	}
	defer stmt.Close()

	err = stmt.QueryRowContext(ctx, email).Scan(&auth.Id, &auth.Username, &auth.Email, &auth.Password, &auth.Role)
	if err == sql.ErrNoRows {
		return &domain.Auth{}, errors.New("No exist user ")
	}
	if err != nil {
		return &domain.Auth{}, err
	}
	return auth, nil
}
