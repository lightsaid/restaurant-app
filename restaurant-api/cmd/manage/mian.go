package main

import (
	"os"

	ut "github.com/go-playground/universal-translator"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
	"lightsaid.com/restaurant-app/restaurant-api/internal/driver"
	"lightsaid.com/restaurant-app/restaurant-api/internal/logger"
	"lightsaid.com/restaurant-app/restaurant-api/internal/repository"
	"lightsaid.com/restaurant-app/restaurant-api/internal/repository/dbrepo"
	"lightsaid.com/restaurant-app/restaurant-api/internal/security"
	"lightsaid.com/restaurant-app/restaurant-api/internal/validator"
)

type application struct {
	port     string
	env      string
	db       *mongo.Client
	trans    ut.Translator
	jwtMaker *security.JWTMaker
	Repo     repository.Repository
}

func main() {
	// 初始化配置
	err := godotenv.Load("manage.dev.env")
	handleFatal(err)

	// 初始化全局日志输出
	l, err := logger.NewLogger(os.Getenv("LOGGER"), "stderr")
	handleFatal(err)
	defer l.Sync()

	// 验证器初始化
	trans, err := validator.InitTrans("zh")
	handleFatal(err)

	// 初始化 JWTMaker 管理token
	jwtMaker, err := security.NewJWTMaker(os.Getenv("JWT_SECRET_KEY"))
	handleFatal(err)

	// 初始化mongo
	client, err := driver.Connect()
	handleFatal(err)
	defer driver.Close(client)

	app := application{
		port:     os.Getenv("PORT"),
		env:      os.Getenv("ENV"),
		db:       client,
		trans:    trans,
		jwtMaker: jwtMaker,
		Repo:     dbrepo.NewMongoRepo(client, "db_restaurant"),
	}
	if err := app.serve(); err != nil {
		zap.S().Fatal(err)
	}
}
