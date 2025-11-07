package routes

import (
	"otochope/database/helpers"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func InitItemsRoutes(r *gin.Engine) {
	itemsGroup := r.Group("/items")

	itemsGroup.GET("/", func(c *gin.Context) {
		items, err := helpers.GetActiveItems()
		ReturnDataOrError(c, items, err)
	})
	itemsGroup.GET("/all", func(c *gin.Context) {
		items, err := helpers.GetAllItems()
		ReturnDataOrError(c, items, err)
	})
	itemsGroup.GET("/statistics", func(c *gin.Context) {
		stats, err := helpers.GetItemsStatistics()
		ReturnDataOrError(c, stats, err)
	})

	/* Single Item Group */

	singleItemGroup := itemsGroup.Group("/:uid", func(c *gin.Context) {
		uidStr := c.Param("uid")
		uid, err := uuid.Parse(uidStr)
		if err != nil {
			c.JSON(400, gin.H{"error": "Invalid UUID"})
			c.Abort()
			return
		}
		c.Set("item_uid", uid)
		c.Next()
	})

	singleItemGroup.GET("/", func(c *gin.Context) {
		item_uid := c.MustGet("item_uid").(uuid.UUID)
		item, err := helpers.GetItemByUID(item_uid)

		ReturnDataOrError(c, item, err)
	})
	singleItemGroup.GET("/inventories", func(c *gin.Context) {
		item_uid := c.MustGet("item_uid").(uuid.UUID)
		inventories, err := helpers.GetActiveUserInventoryItemsByItemUID(item_uid)

		ReturnDataOrError(c, inventories, err)
	})
	singleItemGroup.GET("/inventories/all", func(c *gin.Context) {
		item_uid := c.MustGet("item_uid").(uuid.UUID)
		inventories, err := helpers.GetAllUserInventoryItemsByItemUID(item_uid)

		ReturnDataOrError(c, inventories, err)
	})
	singleItemGroup.GET("/carts", func(c *gin.Context) {
		item_uid := c.MustGet("item_uid").(uuid.UUID)
		carts, err := helpers.GetActiveCartItemsByItemUID(item_uid)

		ReturnDataOrError(c, carts, err)
	})
	singleItemGroup.GET("/carts/all", func(c *gin.Context) {
		item_uid := c.MustGet("item_uid").(uuid.UUID)
		carts, err := helpers.GetAllCartItemsByItemUID(item_uid)

		ReturnDataOrError(c, carts, err)
	})
}
