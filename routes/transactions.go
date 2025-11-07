package routes

import (
	"otochope/database/helpers"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func InitTransactionsRoutes(r *gin.Engine) {
	transactionsGroup := r.Group("/transactions")

	transactionsGroup.GET("/", func(c *gin.Context) {
		transactions, err := helpers.GetAllTransactions()
		ReturnDataOrError(c, transactions, err)
	})
	transactionsGroup.GET("/statistics", func(c *gin.Context) {
		stats, err := helpers.GetTransactionsStatistics()
		ReturnDataOrError(c, stats, err)
	})
	transactionsGroup.GET("/status/:status", func(c *gin.Context) {
		status := c.Param("status")
		transactions, err := helpers.GetAllTransactionsByStatus(status)
		ReturnDataOrError(c, transactions, err)
	})

	transactionGroup := transactionsGroup.Group("/:uid", func(c *gin.Context) {
		uidStr := c.Param("uid")
		uid, err := uuid.Parse(uidStr)
		if err != nil {
			c.JSON(400, gin.H{"error": "Invalid UUID"})
			c.Abort()
			return
		}
		c.Set("transaction_uid", uid)
		c.Next()
	})

	transactionGroup.GET("/", func(c *gin.Context) {
		uid := c.MustGet("transaction_uid").(uuid.UUID)
		transaction, err := helpers.GetTransactionByUID(uid)
		ReturnDataOrError(c, transaction, err)
	})
	transactionGroup.GET("/user", func(c *gin.Context) {
		uid := c.MustGet("transaction_uid").(uuid.UUID)
		transaction, err := helpers.GetTransactionByUID(uid)
		if err != nil {
			ReturnDataOrError(c, nil, err)
			return
		}

		user, err := transaction.GetUser()
		ReturnDataOrError(c, user, err)
	})
	transactionGroup.GET("/cart", func(c *gin.Context) {
		uid := c.MustGet("transaction_uid").(uuid.UUID)
		transaction, err := helpers.GetTransactionByUID(uid)
		if err != nil {
			ReturnDataOrError(c, nil, err)
			return
		}

		cart, err := transaction.GetCart()
		ReturnDataOrError(c, cart, err)
	})
}
