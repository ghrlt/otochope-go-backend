package helpers

import (
	"otochope/database"

	"github.com/google/uuid"
)

func GetAllCarts() ([]Cart, error) {
	query := "SELECT * FROM user_carts"
	rows, err := database.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return parseCartRows(rows)
}

func GetActiveCarts() ([]Cart, error) {
	query := "SELECT * FROM user_carts WHERE active = 1"
	rows, err := database.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return parseCartRows(rows)
}

func GetCartByUID(uid uuid.UUID) (*Cart, error) {
	query := "SELECT * FROM user_carts WHERE uid = ?"
	rows, err := database.Query(query, uid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return parseCartRow(rows)
}

func GetAllCartsByUserUID(userUID uuid.UUID) ([]Cart, error) {
	query := "SELECT * FROM user_carts WHERE user_uid = ?"
	rows, err := database.Query(query, userUID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return parseCartRows(rows)
}

func GetActiveCartByUserUID(userUID uuid.UUID) (*Cart, error) {
	query := "SELECT * FROM user_carts WHERE user_uid = ? AND active = 1 LIMIT 1"
	rows, err := database.Query(query, userUID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return parseCartRow(rows)
}
