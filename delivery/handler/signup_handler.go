package handler

import (
	"lynx/domain"

	"github.com/gofiber/fiber/v2"
)

type signupHandler struct {
	signupUsecase domain.SignupUsecase
}

func NewSignupHandler(signupUsecase domain.SignupUsecase) *signupHandler {
	return &signupHandler{
		signupUsecase: signupUsecase,
	}
}

func (sh *signupHandler) Signup(c *fiber.Ctx) error {
	signupReq := new(domain.SignupRequest)

	if err := c.BodyParser(signupReq); err != nil {
		return FiberError(c, fiber.NewError(fiber.StatusBadRequest, err.Error()))
	}

	if err := validateReq(signupReq); err != nil {
		return FiberError(c, err)
	}

	resp, err := sh.signupUsecase.Signup(signupReq)
	if err != nil {
		return FiberError(c, err)
	}

	return c.JSON(resp)
}
