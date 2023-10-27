package usecase

import (
	"lynx/bootstrap"
	"lynx/domain"
	"lynx/internal/logger"
	"lynx/internal/tokenutil"
	"lynx/model"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type signupUsecase struct {
	userRepository domain.UserRepository
	env            *bootstrap.Env
}

func NewSignupUsecase(userRepository domain.UserRepository, env *bootstrap.Env) domain.SignupUsecase {
	return &signupUsecase{
		userRepository: userRepository,
		env:            env,
	}
}

func (su *signupUsecase) Signup(signupReq *domain.SignupRequest) (signupResp domain.SignupResponse, err error) {
	_, err = su.userRepository.GetByEmail(signupReq.Email)
	if err != nil && err != mongo.ErrNoDocuments {
		logger.Error(err)
		return
	}

	if err == nil {
		return domain.SignupResponse{}, domain.ErrEmailExist
	}

	bytesHash, err := bcrypt.GenerateFromPassword([]byte(signupReq.Password), 10)
	if err != nil {
		logger.Error(err)
		return
	}

	now := time.Now()
	userForSave := &model.User{
		UserID:    uuid.NewString(),
		Provider:  model.LocalProvider,
		Email:     signupReq.Email,
		Password:  string(bytesHash),
		Firstname: signupReq.Firstname,
		Lastname:  signupReq.Lastname,
		Avatar:    signupReq.Avatar,
		Role:      model.UserRole,
		IsActive:  true,
		CreatedAt: now,
		UpdatedAt: now,
	}

	err = su.userRepository.Create(userForSave)
	if err != nil {
		logger.Error(err)
		return
	}

	accessTokenString, err := tokenutil.NewAccessClaims(userForSave, 2).JwtString(su.env.Jwt.AccessSecret)
	if err != nil {
		logger.Error(err)
		return
	}

	refreshTokenString, err := tokenutil.NewRefreshClaims(userForSave, 24).JwtString(su.env.Jwt.RefreshSecret)
	if err != nil {
		logger.Error(err)
		return
	}

	signupResp = domain.SignupResponse{
		AccessToken:  accessTokenString,
		RefreshToken: refreshTokenString,
	}
	return
}
