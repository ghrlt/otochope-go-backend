package helpers

import (
	"otochope/database"

	"github.com/google/uuid"
)

func GetAllCartItemsByCartUID(cartUID uuid.UUID) ([]CartItem, error) {
	query := "SELECT * FROM user_cart_items WHERE cart_uid = ?"
	rows, err := database.Query(query, cartUID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return parseCartItemRows(rows)
}

func GetActiveCartItemsByCartUID(cartUID uuid.UUID) ([]CartItem, error) {
	query := "SELECT * FROM user_cart_items WHERE cart_uid = ? AND active = 1"
	rows, err := database.Query(query, cartUID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return parseCartItemRows(rows)
}

func GetAllCartItemsByItemUID(itemUID uuid.UUID) ([]CartItem, error) {
	query := "SELECT * FROM user_cart_items WHERE item_uid = ?"
	rows, err := database.Query(query, itemUID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return parseCartItemRows(rows)
}

func GetActiveCartItemsByItemUID(itemUID uuid.UUID) ([]CartItem, error) {
	query := "SELECT * FROM user_cart_items WHERE item_uid = ? AND active = 1"
	rows, err := database.Query(query, itemUID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return parseCartItemRows(rows)
}

func GetCartItemByUID(uid uuid.UUID) (*CartItem, error) {
	query := "SELECT * FROM user_cart_items WHERE uid = ? LIMIT 1"
	rows, err := database.Query(query, uid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return parseCartItemRow(rows)
}

func GetActiveCartItemByUID(uid uuid.UUID) (*CartItem, error) {
	query := "SELECT * FROM user_cart_items WHERE uid = ? AND active = 1 LIMIT 1"
	rows, err := database.Query(query, uid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return parseCartItemRow(rows)
}
