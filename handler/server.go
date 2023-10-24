package handler

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"gitlab.com/mandoor/api/controllers"
	"gitlab.com/mandoor/api/pb"
	"gitlab.com/mandoor/bootstrap"
	"gitlab.com/mandoor/internal/token"
	"gitlab.com/mandoor/repository"
	"gitlab.com/mandoor/usecase"
)

type Server struct {
	pb.UnimplementedMandoorServer
	config         *bootstrap.Env
	mysqlClient    bootstrap.MysqlClient
	tokenMaker     token.Maker
	router         *gin.Engine
	userController controllers.IUserController
}

func NewServer(app bootstrap.Application) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(app.Env.TOKEN_KEY)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	repoUser := repository.NewUserRepository(app.Mysql)
	userAdminUseCase := usecase.NewUseCaseUser(repoUser)
	ctrlUser := controllers.NewUserController(app.Validator, userAdminUseCase)

	server := &Server{
		config:         app.Env,
		mysqlClient:    app.Mysql,
		tokenMaker:     tokenMaker,
		userController: ctrlUser,
	}

	return server, nil
}

func (s *Server) LoginUser(ctx context.Context, pb *pb.LoginUserRequest) (*pb.LoginUserResponse, error) {
	s.userController.LoginUser(ctx, pb)
	return nil, nil
}

func (s *Server) CreateUser(context.Context, *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {

	return nil, nil
}
