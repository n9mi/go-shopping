package service

import (
	"context"
	"go-shopping/internal/entity"
	"go-shopping/internal/model"
	"go-shopping/internal/repository"
	"go-shopping/internal/utils"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type AuthServiceImpl struct {
	ViperCfg       *viper.Viper
	Logger         *logrus.Logger
	DB             *gorm.DB
	Validator      *validator.Validate
	UserRepository *repository.UserRepository
}

func NewAuthService(viperCfg *viper.Viper, logger *logrus.Logger, db *gorm.DB,
	validator *validator.Validate, userRepository *repository.UserRepository) AuthService {
	return &AuthServiceImpl{
		ViperCfg:       viperCfg,
		Logger:         logger,
		DB:             db,
		Validator:      validator,
		UserRepository: userRepository,
	}
}

func (s *AuthServiceImpl) Register(ctx context.Context, request *model.RegisterRequest) error {
	tx := s.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := s.Validator.Struct(request); err != nil {
		s.Logger.Warnf("invalid request body : %+v", err)
		return err
	}

	if emailExists := s.UserRepository.ExistsByEmail(tx, request.Email); emailExists {
		s.Logger.Warnf("email %s already used", request.Email)
		return fiber.NewError(fiber.StatusConflict, "email already used")
	}

	userPassword, err := utils.GeneratePassword(request.Password)
	if err != nil {
		s.Logger.Warnf("can't generate user password : %+v", err)
		return err
	}
	newUser := &entity.User{
		ID:       uuid.NewString(),
		Name:     request.Name,
		Email:    request.Email,
		Password: userPassword,
	}
	if err := s.UserRepository.Repository.Create(tx, newUser); err != nil {
		s.Logger.Warnf("failed to create new user : %+v", err)
		return err
	}
	if err := tx.Commit().Error; err != nil {
		s.Logger.Warnf("failed to commit transaction : %+v", err)
		return err
	}

	return nil
}
