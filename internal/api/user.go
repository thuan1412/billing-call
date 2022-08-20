package api

import (
	entCall "calling-bill/ent/call"
	entUser "calling-bill/ent/user"
	"calling-bill/helpers"
	"entgo.io/ent/dialect/sql"
	"github.com/gin-gonic/gin"
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
	var billingData []BillingData
	var err error
	username := c.Param("username")

	// check user existent
	userId, err := helpers.DbClient.User.Query().Where(entUser.Username(username)).FirstID(c)

	if err != nil {
		c.JSON(200, gin.H{
			"asd": err.Error(),
		})
		return
	}

	err = helpers.DbClient.Debug().Call.
		Query().
		Where(entCall.HasUserWith(entUser.ID(userId))).
		Modify(func(s *sql.Selector) {
			s.Select(
				sql.As(sql.Count("*"), "call_count"),
				sql.As(sql.Sum("block_count"), "block_count"),
			)
		}).
		Scan(c, &billingData)

	if err != nil {
		return
	}
	if len(billingData) != 1 {
		//return errors.New("some error :D")
		return
	}

	c.JSON(200, gin.H{
		"data": billingData[0],
	})
	return
}
