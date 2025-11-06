package helpers

import (
	"otochope/database"

	"github.com/google/uuid"
)

func GetAllItems() ([]Item, error) {
	query := "SELECT * FROM items"
	rows, err := database.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return parseItemRows(rows)
}

func GetActiveItems() ([]Item, error) {
	query := "SELECT * FROM items WHERE active = 1"
	rows, err := database.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return parseItemRows(rows)
}

func GetAllItemsByCategoryUID(categoryUID uuid.UUID) ([]Item, error) {
	query := "SELECT * FROM items WHERE category_uid = ?"
	rows, err := database.Query(query, categoryUID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return parseItemRows(rows)
}

func GetActiveItemsByCategoryUID(categoryUID uuid.UUID) ([]Item, error) {
	query := "SELECT * FROM items WHERE category_uid = ? AND active = 1"
	rows, err := database.Query(query, categoryUID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return parseItemRows(rows)
}

func GetItemByUID(uid uuid.UUID) (*Item, error) {
	query := "SELECT * FROM items WHERE uid = ?"
	rows, err := database.Query(query, uid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return parseItemRow(rows)
}

func GetActiveItemByUID(uid uuid.UUID) (*Item, error) {
	query := "SELECT * FROM items WHERE uid = ? AND active = 1"
	rows, err := database.Query(query, uid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return parseItemRow(rows)
}
