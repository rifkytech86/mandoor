package usecase

import (
	"context"
	"gitlab.com/mandoor/repository"
)

type useCaseUserAdmin struct {
	UserRepository repository.IUserRepository
}

//go:generate mockery --name IUseCaseUserAdmin
type IUseCaseUser interface {
	Login(context.Context) (string, error)
}

func NewUseCaseUser(repoUser repository.IUserRepository) IUseCaseUser {
	return &useCaseUserAdmin{
		UserRepository: repoUser,
	}
}

func (u *useCaseUserAdmin) Login(context.Context) (string, error) {

	return "", nil
}
