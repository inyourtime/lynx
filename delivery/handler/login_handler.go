package handler

import (
	"lynx/domain"

	"github.com/gofiber/fiber/v2"
)

type loginHandler struct {
	loginUsecase domain.LoginUsecase
}

func NewLoginHandler(loginUsecase domain.LoginUsecase) *loginHandler {
	return &loginHandler{
		loginUsecase: loginUsecase,
	}
}

func (lh *loginHandler) Login(c *fiber.Ctx) error {
	loginReq := new(domain.LoginRequest)

	if err := c.BodyParser(loginReq); err != nil {
		return FiberError(c, fiber.NewError(fiber.StatusBadRequest, err.Error()))
	}

	if err := validateReq(loginReq); err != nil {
		return FiberError(c, err)
	}

	resp, err := lh.loginUsecase.Login(loginReq)
	if err != nil {
		return FiberError(c, err)
	}

	return c.JSON(resp)
}
