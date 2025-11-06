package helpers

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Item struct {
	UID             uuid.UUID
	CategoryUID     uuid.UUID
	Name            string
	Description     string
	PriceInEURCents int64
	CreatedAt       time.Time
	Active          bool
}

func (i *Item) GetCategory() (*ItemCategory, error) {
	return GetItemCategoryByUID(i.CategoryUID)
}

func parseItemRows(rows *sql.Rows) ([]Item, error) {
	var items []Item
	for rows.Next() {
		var item Item
		if err := rows.Scan(&item.UID, &item.CategoryUID, &item.Name, &item.Description, &item.PriceInEURCents, &item.CreatedAt, &item.Active); err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, nil
}

func parseItemRow(rows *sql.Rows) (*Item, error) {
	var item Item
	if !rows.Next() {
		return nil, nil // No item found
	}
	if err := rows.Scan(&item.UID, &item.CategoryUID, &item.Name, &item.Description, &item.PriceInEURCents, &item.CreatedAt, &item.Active); err != nil {
		return nil, err
	}
	return &item, nil
}
