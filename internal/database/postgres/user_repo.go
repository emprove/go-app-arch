package postgres

import (
	"context"
	"encoding/json"

	"go-app-arch/internal/database"
	"go-app-arch/internal/entity"
	"go-app-arch/internal/repository"

	"github.com/jackc/pgx/v5"
)

type userRepo struct {
	db database.DB
}

func NewUserRepository(db database.DB) repository.User {
	return &userRepo{db: db}
}

func (repo *userRepo) FindOneByToken(token string) (*entity.User, error) {
	type dbrow struct {
		ID          int
		Name        string
		Email       string
		Permissions string
	}
	query := `
		SELECT  
			u.id as "id", 
			u.name as "name", 
			u.email as "email", 
			u.permissions as "permissions"
		FROM users u
		WHERE u.access_token = $1`

	rows, err := repo.db.Query(context.TODO(), query, token)
	if err != nil {
		return nil, err
	}

	m, err := pgx.CollectExactlyOneRow(rows.(pgx.Rows), pgx.RowToStructByName[dbrow])
	if err != nil {
		return nil, err
	}

	var res entity.User
	res.ID = m.ID
	res.Name = m.Name
	res.Email = m.Email
	if err := json.Unmarshal([]byte(m.Permissions), &res.Permissions); err != nil {
		return nil, err
	}

	return &res, nil
}
