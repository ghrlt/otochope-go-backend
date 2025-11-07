package routes

import (
	"otochope/database/helpers"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func InitCouponsRoutes(r *gin.Engine) {
	couponsGroup := r.Group("/coupons")

	couponsGroup.GET("/", func(c *gin.Context) {
		coupons, err := helpers.GetActiveCoupons()
		ReturnDataOrError(c, coupons, err)
	})
	couponsGroup.GET("/all", func(c *gin.Context) {
		coupons, err := helpers.GetAllCoupons()
		ReturnDataOrError(c, coupons, err)
	})
	couponsGroup.GET("/statistics", func(c *gin.Context) {
		stats, err := helpers.GetCouponsStatistics()
		ReturnDataOrError(c, stats, err)
	})

	couponGroup := couponsGroup.Group("/:uid", func(c *gin.Context) {
		uidStr := c.Param("uid")
		uid, err := uuid.Parse(uidStr)
		if err != nil {
			c.JSON(400, gin.H{"error": "Invalid UUID"})
			c.Abort()
			return
		}
		c.Set("coupon_uid", uid)
		c.Next()
	})

	couponGroup.GET("/", func(c *gin.Context) {
		coupon_uid := c.MustGet("coupon_uid").(uuid.UUID)
		coupon, err := helpers.GetCouponByUID(coupon_uid)

		ReturnDataOrError(c, coupon, err)
	})
	couponGroup.GET("/carts", func(c *gin.Context) {
		coupon_uid := c.MustGet("coupon_uid").(uuid.UUID)
		carts, err := helpers.GetActiveCartCouponsByCouponUID(coupon_uid)

		ReturnDataOrError(c, carts, err)
	})
	couponGroup.GET("/carts/all", func(c *gin.Context) {
		coupon_uid := c.MustGet("coupon_uid").(uuid.UUID)
		carts, err := helpers.GetAllCartCouponsByCouponUID(coupon_uid)

		ReturnDataOrError(c, carts, err)
	})
}
