package controllers

import (
	"context"
	"gitlab.com/mandoor/api/pb"
	"gitlab.com/mandoor/internal/validator"
	"gitlab.com/mandoor/usecase"
)

//go:generate mockery --name IUserController
type IUserController interface {
	LoginUser(context.Context, *pb.LoginUserRequest) (*pb.LoginUserResponse, error)
}

type userController struct {
	UserUseCase usecase.IUseCaseUser
	Validator   validator.IValidator
}

func NewUserController(validator validator.IValidator, userUseCase usecase.IUseCaseUser) IUserController {
	return &userController{
		UserUseCase: userUseCase,
		Validator:   validator,
	}
}

func (u *userController) LoginUser(ctx context.Context, req *pb.LoginUserRequest) (*pb.LoginUserResponse, error) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}

	return nil, nil
}
