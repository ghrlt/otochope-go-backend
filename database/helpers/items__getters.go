package helpers

import (
	"otochope/database"

	"github.com/google/uuid"
)

type ItemStats struct {
	TotalItems       int
	ActiveItems      int
	InactiveItems    int
	TotalPerCategory map[uuid.UUID]int
}

func GetItemsStatistics() (*ItemStats, error) {
	query := `
		SELECT
			COUNT(*) AS total,
			SUM(CASE WHEN active = 1 THEN 1 ELSE 0 END) AS active,
			SUM(CASE WHEN active = 0 THEN 1 ELSE 0 END) AS inactive
		FROM items;
	`

	rows, err := database.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var stats ItemStats
	if rows.Next() {
		if err := rows.Scan(&stats.TotalItems, &stats.ActiveItems, &stats.InactiveItems); err != nil {
			return nil, err
		}
	}

	// Get total items per category
	categoryQuery := `
		SELECT category_uid, COUNT(*) AS total
		FROM items
		GROUP BY category_uid;
	`
	categoryRows, err := database.Query(categoryQuery)
	if err != nil {
		return nil, err
	}
	defer categoryRows.Close()

	stats.TotalPerCategory = make(map[uuid.UUID]int)
	for categoryRows.Next() {
		var categoryUID uuid.UUID
		var total int
		if err := categoryRows.Scan(&categoryUID, &total); err != nil {
			return nil, err
		}
		stats.TotalPerCategory[categoryUID] = total
	}

	return &stats, nil
}

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

func GetAllItemsByItemCategoryUID(categoryUID uuid.UUID) ([]Item, error) {
	query := "SELECT * FROM items WHERE category_uid = ?"
	rows, err := database.Query(query, categoryUID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return parseItemRows(rows)
}

func GetActiveItemsByItemCategoryUID(categoryUID uuid.UUID) ([]Item, error) {
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
