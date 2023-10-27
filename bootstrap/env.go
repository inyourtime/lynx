package bootstrap

import (
	"log"
	"lynx/internal/logger"
	"os"

	"github.com/joho/godotenv"
)

type Env struct {
	AppEnv         string
	AppPort        string
	ApiPrefix      string
	DiscordWebhook DiscordWebhook
	Db             Db
	S3             S3
	Jwt            Jwt
	Google         Google
}

func NewEnv() *Env {
	if err := godotenv.Load(); err != nil && !os.IsNotExist(err) {
		log.Fatalln("Error loading .env")
	}

	env := Env{
		AppEnv:    os.Getenv("APP_ENV"),
		AppPort:   os.Getenv("PORT"),
		ApiPrefix: os.Getenv("API_PREFIX"),
		DiscordWebhook: DiscordWebhook{
			ID:    os.Getenv("DISCORD_ID"),
			Token: os.Getenv("DISCORD_TOKEN"),
		},
		Db: Db{
			Mongo: Mongo{
				Uri:      os.Getenv("MONGODB_URL"),
				Database: os.Getenv("MONGODB_DATABASE"),
			},
		},
		S3: S3{
			AccountID:       os.Getenv("OBJECTSTORAGE_ACCOUNTID"),
			AccessKeyID:     os.Getenv("OBJECTSTORAGE_ACCESSKEYID"),
			AccessKeySecret: os.Getenv("OBJECTSTORAGE_SECRETACCESSKEY"),
			Bucket:          os.Getenv("OBJECTSTORAGE_BUCKET"),
		},
		Jwt: Jwt{
			AccessSecret:  os.Getenv("JWT_ACCESS_SECRET"),
			RefreshSecret: os.Getenv("JWT_REFRESH_SECRET"),
		},
		Google: Google{
			ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
			ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
			RedirectUrl:  os.Getenv("GOOGLE_REDIRECT_URL"),
		},
	}

	if env.AppEnv == "development" {
		logger.Info("The App is running in development env")
	}

	return &env
}

type DiscordWebhook struct {
	ID    string
	Token string
}

type Db struct {
	Mongo Mongo
}

type Mongo struct {
	Uri      string
	Database string
}

type S3 struct {
	AccountID       string
	AccessKeyID     string
	AccessKeySecret string
	Bucket          string
}

type Jwt struct {
	AccessSecret  string
	RefreshSecret string
}

type Google struct {
	ClientID     string
	ClientSecret string
	RedirectUrl  string
}
