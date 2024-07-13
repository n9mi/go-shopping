package service

import (
	"context"
	"go-shopping/internal/model"
)

type AuthService interface {
	Register(ctx context.Context, request *model.RegisterRequest) error
}
