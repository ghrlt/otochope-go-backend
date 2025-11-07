package routes

import (
	"otochope/utilities"

	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to OTOCHOPE!",
		})
	})

	InitUsersRoutes(r)
	InitTransactionsRoutes(r)
	InitInventoryRoutes(r)
	InitCartsRoutes(r)
	InitCouponsRoutes(r)
	InitItemsRoutes(r)
	InitItemCategoriesRoutes(r)
}

func ReturnDataOrError(c *gin.Context, data any, err error) {
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	if utilities.IsNilOrEmpty(data) {
		c.JSON(204, gin.H{"message": "No data found"})
		return
	}

	c.JSON(200, data)
}
