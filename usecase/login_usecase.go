package usecase

import (
	"lynx/bootstrap"
	"lynx/domain"
	"lynx/internal/logger"
	"lynx/internal/tokenutil"

	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type loginUsecase struct {
	userRepository domain.UserRepository
	env            *bootstrap.Env
}

func NewLoginUsecase(userRepository domain.UserRepository, env *bootstrap.Env) domain.LoginUsecase {
	return &loginUsecase{
		userRepository: userRepository,
		env:            env,
	}
}

func (lu *loginUsecase) Login(loginReq *domain.LoginRequest) (loginResp domain.LoginResponse, err error) {
	user, err := lu.userRepository.GetByEmail(loginReq.Email)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return domain.LoginResponse{}, domain.ErrEmailPwdIncorrect
		}
		logger.Error(err)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginReq.Password))
	if err != nil {
		return domain.LoginResponse{}, domain.ErrEmailPwdIncorrect
	}

	accessTokenString, err := tokenutil.NewAccessClaims(&user, 2).JwtString(lu.env.Jwt.AccessSecret)
	if err != nil {
		logger.Error(err)
		return
	}

	refreshTokenString, err := tokenutil.NewRefreshClaims(&user, 24).JwtString(lu.env.Jwt.RefreshSecret)
	if err != nil {
		logger.Error(err)
		return
	}

	loginResp = domain.LoginResponse{
		AccessToken:  accessTokenString,
		RefreshToken: refreshTokenString,
	}
	return
}
