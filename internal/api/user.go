package api

import (
	"calling-bill/ent"
	"calling-bill/helpers"
	"calling-bill/internal/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func AddCall(c *gin.Context) {
	newCall := helpers.DbClient.Call.
		Create().
		SetDuration(1000).
		SetBlockCount(123).
		SetUserID(1)

	_, err := newCall.Save(c)
	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "something?",
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

	c.JSON(200, gin.H{
		"data": billingData,
	})
	return
}
