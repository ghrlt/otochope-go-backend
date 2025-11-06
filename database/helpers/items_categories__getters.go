package helpers

import (
	"otochope/database"

	"github.com/google/uuid"
)

func GetAllItemCategories() ([]ItemCategory, error) {
	query := "SELECT * FROM items_categories"
	rows, err := database.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return parseItemCategoryRows(rows)
}

func GetActiveItemCategories() ([]ItemCategory, error) {
	query := "SELECT * FROM items_categories WHERE active = 1"
	rows, err := database.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return parseItemCategoryRows(rows)
}

func GetAllItemCategoriesByParentUID(parentUID uuid.UUID) ([]ItemCategory, error) {
	query := "SELECT * FROM items_categories WHERE parent_item_category_uid = ?"
	rows, err := database.Query(query, parentUID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return parseItemCategoryRows(rows)
}

func GetActiveItemCategoriesByParentUID(parentUID uuid.UUID) ([]ItemCategory, error) {
	query := "SELECT * FROM items_categories WHERE parent_item_category_uid = ? AND active = 1"
	rows, err := database.Query(query, parentUID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return parseItemCategoryRows(rows)
}

func GetItemCategoryByUID(uid uuid.UUID) (*ItemCategory, error) {
	query := "SELECT * FROM items_categories WHERE uid = ?"
	rows, err := database.Query(query, uid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return parseItemCategoryRow(rows)
}

func GetActiveItemCategoryByUID(uid uuid.UUID) (*ItemCategory, error) {
	query := "SELECT * FROM items_categories WHERE uid = ? AND active = 1"
	rows, err := database.Query(query, uid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return parseItemCategoryRow(rows)
}
