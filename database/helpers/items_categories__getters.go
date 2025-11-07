package helpers

import (
	"otochope/database"

	"github.com/google/uuid"
)

type ItemCategoryStats struct {
	TotalItemCategories    int
	ActiveItemCategories   int
	InactiveItemCategories int
}

func GetItemCategoriesStatistics() (*ItemCategoryStats, error) {
	query := `
		SELECT
			COUNT(*) AS total,
			SUM(CASE WHEN active = 1 THEN 1 ELSE 0 END) AS active,
			SUM(CASE WHEN active = 0 THEN 1 ELSE 0 END) AS inactive
		FROM item_categories;
	`

	rows, err := database.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var stats ItemCategoryStats
	if rows.Next() {
		if err := rows.Scan(&stats.TotalItemCategories, &stats.ActiveItemCategories, &stats.InactiveItemCategories); err != nil {
			return nil, err
		}
	}

	return &stats, nil
}

func GetAllItemCategories() ([]ItemCategory, error) {
	query := "SELECT * FROM item_categories"
	rows, err := database.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return parseItemCategoryRows(rows)
}

func GetActiveItemCategories() ([]ItemCategory, error) {
	query := "SELECT * FROM item_categories WHERE active = 1"
	rows, err := database.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return parseItemCategoryRows(rows)
}

func GetAllItemCategoriesByParentUID(parentUID uuid.UUID) ([]ItemCategory, error) {
	query := "SELECT * FROM item_categories WHERE parent_item_category_uid = ?"
	rows, err := database.Query(query, parentUID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return parseItemCategoryRows(rows)
}

func GetActiveItemCategoriesByParentUID(parentUID uuid.UUID) ([]ItemCategory, error) {
	query := "SELECT * FROM item_categories WHERE parent_item_category_uid = ? AND active = 1"
	rows, err := database.Query(query, parentUID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return parseItemCategoryRows(rows)
}

func GetItemCategoryByUID(uid uuid.UUID) (*ItemCategory, error) {
	query := "SELECT * FROM item_categories WHERE uid = ?"
	rows, err := database.Query(query, uid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return parseItemCategoryRow(rows)
}

func GetActiveItemCategoryByUID(uid uuid.UUID) (*ItemCategory, error) {
	query := "SELECT * FROM item_categories WHERE uid = ? AND active = 1"
	rows, err := database.Query(query, uid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return parseItemCategoryRow(rows)
}
