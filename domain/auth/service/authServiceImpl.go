package service

import (
	"errors"
	"go-crud/domain/auth/repository"
	"go-crud/entities"
	"go-crud/errs"
	"go-crud/models"
	"go-crud/utils"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/go-sql-driver/mysql"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

type authServiceImpl struct {
	authUserRepository repository.AuthUserRepository
	validate           *validator.Validate
}

func NewAuthService(
	authUserRepository repository.AuthUserRepository,
) AuthService {
	return &authServiceImpl{
		authUserRepository: authUserRepository,
		validate:           validator.New(validator.WithRequiredStructEnabled()),
	}
}

func (s *authServiceImpl) Login(loginRequest *models.LoginRequest) (*models.LoginResponse, error) {

	if err := s.validate.Struct(loginRequest); err != nil {
		return nil, errs.ParseValidationErrors(err)
	}

	user, err := s.authUserRepository.FindByUsername(loginRequest.Username)
	if err != nil {
		zap.L().Error(err.Error())
		return nil, errs.NewNotFoundError("user not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginRequest.Password))
	if err != nil {
		return nil, errs.NewBadRequestError("invalid password")
	}

	accessToken, err := utils.GenerateAccessToken(user.ID, time.Now().Add(24*time.Hour))
	if err != nil {
		zap.L().Error(err.Error())
		return nil, errs.NewInternalError()
	}

	return &models.LoginResponse{
		AccessToken: *accessToken,
	}, nil

}

func (s *authServiceImpl) Register(req *models.RegisterRequest) error {
	if err := s.validate.Struct(req); err != nil {
		return errs.ParseValidationErrors(err)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		zap.L().Error("hash password failed", zap.Error(err))
		return errs.NewInternalError()
	}

	user := &entities.User{
		Username: req.Username,
		Email:    req.Email,
		Password: string(hashedPassword),
	}

	if _, err := s.authUserRepository.Insert(user); err != nil {

		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
			return errs.NewBadRequestError("username or email already exists")
		}

		zap.L().Error("insert user failed", zap.Error(err))
		return errs.NewInternalError()
	}

	return nil
}
