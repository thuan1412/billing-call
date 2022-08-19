package main

import (
	"calling-bill/ent"
	"calling-bill/helpers"
	"calling-bill/route"
	"github.com/gin-gonic/gin"

	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

var logger *zap.Logger
var entClient *ent.Client
var err error

func init() {
	logger = zap.NewExample()
	entClient, err = helpers.GetDb()
	if err != nil {
		panic(err)
	}
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

// panicErr panics if error not nil
func panicErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	app := gin.Default()

	// middlewares
	app.Use(InjectDbClient())
	app.Use(InjectLogger())

	route.SetUp(app)
	err := app.Run() // TODO: read port from env
	panic(err)
}
