package repository

import (
	"context"
	"gitlab.com/mandoor/bootstrap"
)

type userRepository struct {
	MysqlClient bootstrap.MysqlClient
}

//go:generate mockery --name IUserRepository
type IUserRepository interface {
	Fetch(ctx context.Context, filter interface{}, result interface{}) error
}

func NewUserRepository(client bootstrap.MysqlClient) IUserRepository {
	return &userRepository{
		MysqlClient: client,
	}
}

func (u *userRepository) Fetch(ctx context.Context, filter interface{}, result interface{}) error {

	return nil
}
