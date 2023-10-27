package route

import (
	"lynx/bootstrap"
	"lynx/delivery/handler"
	"lynx/repository"
	"lynx/usecase"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewAuthRouter(env *bootstrap.Env, db *mongo.Database, router fiber.Router) {
	authRouter := router.Group("/auth")

	userRepository := repository.NewUserRepository(db)
	loginUsecase := usecase.NewLoginUsecase(userRepository, env)
	loginHandler := handler.NewLoginHandler(loginUsecase)
	signupUsecase := usecase.NewSignupUsecase(userRepository, env)
	signupHandler := handler.NewSignupHandler(signupUsecase)

	{
		authRouter.Post("/login", loginHandler.Login)
		authRouter.Post("/signup", signupHandler.Signup)
	}
}
