package main

import (
	"calling-bill/ent"
	helpers "calling-bill/internal/helpers"
	"calling-bill/route"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

var logger *zap.Logger
var entClient *ent.Client
var err error

func init() {
	err := godotenv.Load()
	helpers.PanicErr(err)
	logger = zap.NewExample()
	entClient, err = helpers.GetDb()
	helpers.PanicErr(err)
}

// InjectDbClient sets dbClient object to the gin context
func InjectDbClient() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("db", entClient)
	}
}

// InjectLogger sets logger object to the gin context
func InjectLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("logger", logger)
	}
}

func main() {
	helpers.MigrateDb(helpers.DbClient)
	if err != nil {
		logger.Info(fmt.Sprintf("error when loading .env %s", err.Error()))
	}
	app := gin.Default()

	// middlewares
	app.Use(InjectDbClient())
	app.Use(InjectLogger())

	route.SetUp(app)
	err = app.Run() // TODO: read port from env
	helpers.PanicErr(err)
}
