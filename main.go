package main

import (
	"fmt"
	"go-crud/database"
	"go-crud/routes"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	defer zap.L().Sync()
	initZapLogger()
	initViperConfig()
	database.InitMysqlConnection()

	fiber := fiber.New()
	routes.RegisterProductRoute(fiber)
	fiber.Listen(fmt.Sprintf(":%v", viper.GetString("server.port")))
}

func initViperConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
}

func initZapLogger() {
	config := zap.NewProductionConfig()
	config.Encoding = "console"
	config.EncoderConfig.TimeKey = "timestamp"
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	config.EncoderConfig.StacktraceKey = ""

	log, err := config.Build()
	if err != nil {
		panic(err)
	}
	zap.ReplaceGlobals(log)
}
