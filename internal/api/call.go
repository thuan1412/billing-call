package api

import (
	"calling-bill/helpers"
	"github.com/gin-gonic/gin"
)

func AddCall(c *gin.Context) {
	newCall := helpers.DbClient.Call.
		Create().
		SetDuration(1000).
		SetBlockCount(123)
	_, err := newCall.Save(c)
	if err != nil {
		c.JSON(200, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "something?",
	})
	return
}
