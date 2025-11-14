package routes

import (
	"otochope/database/helpers"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func InitUsersRoutes(r *gin.Engine) {
	usersGroup := r.Group("/users")

	usersGroup.GET("/", func(c *gin.Context) {
		users, err := helpers.GetActiveUsers()
		ReturnDataOrError(c, users, err)
	})
	usersGroup.GET("/all", func(c *gin.Context) {
		users, err := helpers.GetAllUsers()
		ReturnDataOrError(c, users, err)
	})
	usersGroup.POST("/search", func(c *gin.Context) {
		type SearchParams struct {
			UID        *string `json:"uid"`
			Username   *string `json:"username"`
			Identifier *string `json:"identifier"`
		}
		var searchParams SearchParams
		if err := c.ShouldBindJSON(&searchParams); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		if searchParams.UID == nil && searchParams.Username == nil && searchParams.Identifier == nil {
			c.JSON(400, gin.H{"error": "At least one search parameter must be provided"})
			return
		}
		if searchParams.UID != nil && (searchParams.Username != nil || searchParams.Identifier != nil) {
			c.JSON(400, gin.H{"error": "When searching by UID, no other parameters should be provided"})
			return
		}

		var users []helpers.User
		var err error
		switch {
		case searchParams.UID != nil:
			uid, parseErr := uuid.Parse(*searchParams.UID)
			if parseErr != nil {
				c.JSON(400, gin.H{"error": "Invalid UUID format"})
				return
			}
			var user *helpers.User
			user, err = helpers.GetUserByUID(uid)
			if user != nil {
				users = []helpers.User{*user}
			}

		case searchParams.Username != nil && searchParams.Identifier != nil:
			users, err = helpers.FindUsersByUsernameAndIdentifier(*searchParams.Username, *searchParams.Identifier)

		case searchParams.Username != nil:
			users, err = helpers.FindUsersByUsername(*searchParams.Username)

		case searchParams.Identifier != nil:
			users, err = helpers.FindUsersByIdentifier(*searchParams.Identifier)
		}
		ReturnDataOrError(c, users, err)
	})
	usersGroup.GET("/statistics", func(c *gin.Context) {
		stats, err := helpers.GetUsersStatistics()
		ReturnDataOrError(c, stats, err)
	})

	/* User Group */

	userGroup := usersGroup.Group("/:uid", func(c *gin.Context) {
		uidStr := c.Param("uid")
		uid, err := uuid.Parse(uidStr)
		if err != nil {
			c.JSON(400, gin.H{"error": "Invalid UUID"})
			return
		}
		c.Set("user_uid", uid)
		c.Next()
	})

	userGroup.GET("/", func(c *gin.Context) {
		user_uid := c.MustGet("user_uid").(uuid.UUID)
		user, err := helpers.GetUserByUID(user_uid)

		ReturnDataOrError(c, user, err)
	})

	userGroup.GET("/inventories", func(c *gin.Context) {
		user_uid := c.MustGet("user_uid").(uuid.UUID)
		inventories, err := helpers.GetAllUserInventoriesByUserUID(user_uid)

		ReturnDataOrError(c, inventories, err)
	})
	userGroup.GET("/inventory", func(c *gin.Context) {
		user_uid := c.MustGet("user_uid").(uuid.UUID)
		inventory, err := helpers.GetActiveUserInventoryByUserUID(user_uid)

		ReturnDataOrError(c, inventory, err)
	})

	userGroup.GET("/carts", func(c *gin.Context) {
		user_uid := c.MustGet("user_uid").(uuid.UUID)
		carts, err := helpers.GetAllCartsByUserUID(user_uid)

		ReturnDataOrError(c, carts, err)
	})
	userGroup.GET("/cart", func(c *gin.Context) {
		user_uid := c.MustGet("user_uid").(uuid.UUID)
		cart, err := helpers.GetActiveCartByUserUID(user_uid)

		ReturnDataOrError(c, cart, err)
	})

	userGroup.GET("/transactions", func(c *gin.Context) {
		user_uid := c.MustGet("user_uid").(uuid.UUID)
		transactions, err := helpers.GetAllTransactionsByUserUID(user_uid)

		ReturnDataOrError(c, transactions, err)
	})
	userGroup.GET("/transactions/:transaction_uid", func(c *gin.Context) {
		user_uid := c.MustGet("user_uid").(uuid.UUID)
		transactionUIDStr := c.Param("transaction_uid")
		transactionUID, err := uuid.Parse(transactionUIDStr)
		if err != nil {
			c.JSON(400, gin.H{"error": "Invalid transaction UUID"})
			return
		}

		transaction, err := helpers.GetTransactionByUID(transactionUID)
		if transaction != nil && transaction.UserUID != user_uid {
			c.JSON(403, gin.H{"error": "Transaction does not belong to the user"})
			return
		}

		ReturnDataOrError(c, transaction, err)
	})
}
