package routes

import (
	"otochope/database/helpers"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func InitInventoryRoutes(r *gin.Engine) {
	inventoryGroup := r.Group("/inventory")

	inventoryGroup.GET("/", func(c *gin.Context) {
		inventories, err := helpers.GetActiveUserInventories()
		ReturnDataOrError(c, inventories, err)
	})
	inventoryGroup.GET("/all", func(c *gin.Context) {
		inventories, err := helpers.GetAllUserInventories()
		ReturnDataOrError(c, inventories, err)
	})

	/* Single Inventory Group */

	singleInventoryGroup := inventoryGroup.Group("/:uid", func(c *gin.Context) {
		uidStr := c.Param("uid")
		uid, err := uuid.Parse(uidStr)
		if err != nil {
			c.JSON(400, gin.H{"error": "Invalid UUID"})
			c.Abort()
			return
		}
		c.Set("inventory_uid", uid)
		c.Next()
	})

	singleInventoryGroup.GET("/", func(c *gin.Context) {
		inventory_uid := c.MustGet("inventory_uid").(uuid.UUID)
		inventory, err := helpers.GetUserInventoryByUID(inventory_uid)

		ReturnDataOrError(c, inventory, err)
	})
	singleInventoryGroup.GET("/user", func(c *gin.Context) {
		inventory_uid := c.MustGet("inventory_uid").(uuid.UUID)
		inventory, err := helpers.GetUserInventoryByUID(inventory_uid)
		if err != nil {
			ReturnDataOrError(c, nil, err)
			return
		}

		user, err := inventory.GetUser()
		ReturnDataOrError(c, user, err)
	})
	singleInventoryGroup.GET("/items", func(c *gin.Context) {
		inventory_uid := c.MustGet("inventory_uid").(uuid.UUID)
		items, err := helpers.GetActiveUserInventoryItemsByInventoryUID(inventory_uid)

		ReturnDataOrError(c, items, err)
	})
	singleInventoryGroup.GET("/items/all", func(c *gin.Context) {
		inventory_uid := c.MustGet("inventory_uid").(uuid.UUID)
		items, err := helpers.GetAllUserInventoryItemsByInventoryUID(inventory_uid)

		ReturnDataOrError(c, items, err)
	})

}
