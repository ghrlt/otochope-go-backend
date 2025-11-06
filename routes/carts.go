package routes

import (
	"otochope/database/helpers"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func InitCartsRoutes(r *gin.Engine) {
	cartsGroup := r.Group("/carts")

	cartsGroup.GET("/", func(c *gin.Context) {
		carts, err := helpers.GetActiveCarts()
		ReturnDataOrError(c, carts, err)
	})
	cartsGroup.GET("/all", func(c *gin.Context) {
		carts, err := helpers.GetAllCarts()
		ReturnDataOrError(c, carts, err)
	})

	/* Single Cart Group */

	singleCartGroup := cartsGroup.Group("/:uid", func(c *gin.Context) {
		uidStr := c.Param("uid")
		uid, err := uuid.Parse(uidStr)
		if err != nil {
			c.JSON(400, gin.H{"error": "Invalid UUID"})
			c.Abort()
			return
		}
		c.Set("cart_uid", uid)
		c.Next()
	})

	singleCartGroup.GET("/", func(c *gin.Context) {
		cart_uid := c.MustGet("cart_uid").(uuid.UUID)
		cart, err := helpers.GetCartByUID(cart_uid)

		ReturnDataOrError(c, cart, err)
	})
	singleCartGroup.GET("/user", func(c *gin.Context) {
		cart_uid := c.MustGet("cart_uid").(uuid.UUID)
		cart, err := helpers.GetCartByUID(cart_uid)
		if err != nil {
			ReturnDataOrError(c, nil, err)
			return
		}

		user, err := cart.GetUser()
		ReturnDataOrError(c, user, err)
	})
	singleCartGroup.GET("/items", func(c *gin.Context) {
		cart_uid := c.MustGet("cart_uid").(uuid.UUID)
		items, err := helpers.GetActiveCartItemsByCartUID(cart_uid)

		ReturnDataOrError(c, items, err)
	})
	singleCartGroup.GET("/items/all", func(c *gin.Context) {
		cart_uid := c.MustGet("cart_uid").(uuid.UUID)
		items, err := helpers.GetAllCartItemsByCartUID(cart_uid)

		ReturnDataOrError(c, items, err)
	})
	singleCartGroup.GET("/coupons", func(c *gin.Context) {
		cart_uid := c.MustGet("cart_uid").(uuid.UUID)
		coupons, err := helpers.GetActiveCartCouponsByCartUID(cart_uid)

		ReturnDataOrError(c, coupons, err)
	})
	singleCartGroup.GET("/coupons/all", func(c *gin.Context) {
		cart_uid := c.MustGet("cart_uid").(uuid.UUID)
		coupons, err := helpers.GetAllCartCouponsByCartUID(cart_uid)

		ReturnDataOrError(c, coupons, err)
	})
	singleCartGroup.GET("/transactions", func(c *gin.Context) {
		cart_uid := c.MustGet("cart_uid").(uuid.UUID)
		transactions, err := helpers.GetAllTransactionsByCartUID(cart_uid)

		ReturnDataOrError(c, transactions, err)
	})
}
