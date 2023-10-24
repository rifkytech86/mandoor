package handler

import (
	"gitlab.com/mandoor/api/controllers"
	"gitlab.com/mandoor/bootstrap"
)

type MyHandler struct {
	Application bootstrap.Application
}

func NewServiceInitial(app bootstrap.Application) MyHandler {
	return MyHandler{
		Application: app,
	}
}

type ServerInterfaceWrapper struct {
	UserAdminHandler controllers.IUserController
}

func (h *MyHandler) UserAdminHandler() controllers.IUserController {
	userAdminController := controllers.NewUserController(nil, nil)
	return userAdminController
}
