package controller

import (
	"go-shopping/internal/service"

	"github.com/sirupsen/logrus"
)

type ControllerSetup struct {
	AuthController *AuthController
}

func Setup(logger *logrus.Logger, serviceSetup *service.ServiceSetup) *ControllerSetup {
	return &ControllerSetup{
		AuthController: NewAuthController(serviceSetup.AuthService, logger),
	}
}
