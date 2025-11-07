package routes

import (
	"otochope/database/helpers"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func InitItemCategoriesRoutes(r *gin.Engine) {
	itemCategoriesGroup := r.Group("/item-categories")
	itemCategoriesGroup.GET("/", func(c *gin.Context) {
		categories, err := helpers.GetActiveItemCategories()
		ReturnDataOrError(c, categories, err)
	})
	itemCategoriesGroup.GET("/all", func(c *gin.Context) {
		categories, err := helpers.GetAllItemCategories()
		ReturnDataOrError(c, categories, err)
	})
	itemCategoriesGroup.GET("/statistics", func(c *gin.Context) {
		stats, err := helpers.GetItemCategoriesStatistics()
		ReturnDataOrError(c, stats, err)
	})

	itemCategoryGroup := itemCategoriesGroup.Group("/:uid", func(c *gin.Context) {
		uidStr := c.Param("uid")
		uid, err := uuid.Parse(uidStr)
		if err != nil {
			c.JSON(400, gin.H{"error": "Invalid UUID"})
			c.Abort()
			return
		}
		c.Set("item_category_uid", uid)
		c.Next()
	})
	itemCategoryGroup.GET("/", func(c *gin.Context) {
		item_category_uid := c.MustGet("item_category_uid").(uuid.UUID)
		category, err := helpers.GetItemCategoryByUID(item_category_uid)

		ReturnDataOrError(c, category, err)
	})
	itemCategoryGroup.GET("/subcategories", func(c *gin.Context) {
		item_category_uid := c.MustGet("item_category_uid").(uuid.UUID)
		subcategories, err := helpers.GetAllItemCategoriesByParentUID(item_category_uid)

		ReturnDataOrError(c, subcategories, err)
	})
	itemCategoryGroup.GET("/items", func(c *gin.Context) {
		item_category_uid := c.MustGet("item_category_uid").(uuid.UUID)
		items, err := helpers.GetActiveItemsByItemCategoryUID(item_category_uid)

		ReturnDataOrError(c, items, err)
	})
	itemCategoryGroup.GET("/items/all", func(c *gin.Context) {
		item_category_uid := c.MustGet("item_category_uid").(uuid.UUID)
		items, err := helpers.GetAllItemsByItemCategoryUID(item_category_uid)

		ReturnDataOrError(c, items, err)
	})
}
