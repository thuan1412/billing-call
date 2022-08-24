package api

import (
	"calling-bill/ent"
	"calling-bill/internal/helpers"
	"calling-bill/internal/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

func AddCall(c *gin.Context) {
	var err error
	userService := service.UserService{
		DB:     helpers.DbClient,
		Logger: zap.NewExample(), // TODO: should use DI for this field
	}
	username := c.Param("username")

	// check user existent
	userId, err := userService.GetUserIDFromUsername(c, username)

	// validate request body
	var createCallData service.CreateCallData
	if err := c.ShouldBindJSON(&createCallData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err != nil {
		if ent.IsNotFound(err) {
			c.JSON(400, gin.H{
				"message": "invalid username",
			})
			return
		}
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	err = userService.AddCall(c, *userId, createCallData)
	if err != nil {
		return
	}

	c.JSON(200, gin.H{
		"message": "success",
	})
	return
}

type BillingData struct {
	Count int `sql:"call_count"`
	Sum   int `sql:"block_count"`
}

func GetBill(c *gin.Context) {
	var err error
	userService := service.UserService{
		DB:     helpers.DbClient,
		Logger: zap.NewExample(), // TODO: should use DI for this field
	}
	username := c.Param("username")

	// check user existent
	userId, err := userService.GetUserIDFromUsername(c, username)

	if err != nil {
		if ent.IsNotFound(err) {
			c.JSON(400, gin.H{
				"message": "invalid username",
			})
			return
		}
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	billingData, err := userService.GetBilling(c, *userId)

	c.JSON(200, billingData)
	return
}
