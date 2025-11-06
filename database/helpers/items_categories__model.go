package helpers

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type ItemCategory struct {
	UID                   uuid.UUID
	ParentItemCategoryUID *uuid.UUID
	Label                 string
	CreatedAt             time.Time
	Active                bool
}

func (ic *ItemCategory) GetParentCategory() (*ItemCategory, error) {
	if ic.ParentItemCategoryUID == nil {
		return nil, nil
	}
	return GetItemCategoryByUID(*ic.ParentItemCategoryUID)
}

func (ic *ItemCategory) GetSubCategories() ([]ItemCategory, error) {
	return GetAllItemCategoriesByParentUID(ic.UID)
}

func parseItemCategoryRows(rows *sql.Rows) ([]ItemCategory, error) {
	var itemCategories []ItemCategory
	for rows.Next() {
		var itemCategory ItemCategory
		var parentUID sql.NullString
		if err := rows.Scan(&itemCategory.UID, &parentUID, &itemCategory.Label, &itemCategory.CreatedAt, &itemCategory.Active); err != nil {
			return nil, err
		}
		if parentUID.Valid {
			uid, err := uuid.Parse(parentUID.String)
			if err != nil {
				return nil, err
			}
			itemCategory.ParentItemCategoryUID = &uid
		} else {
			itemCategory.ParentItemCategoryUID = nil
		}
		itemCategories = append(itemCategories, itemCategory)
	}
	return itemCategories, nil
}

func parseItemCategoryRow(rows *sql.Rows) (*ItemCategory, error) {
	var itemCategory ItemCategory
	var parentUID sql.NullString
	if !rows.Next() {
		return nil, nil // No item category found
	}
	if err := rows.Scan(&itemCategory.UID, &parentUID, &itemCategory.Label, &itemCategory.CreatedAt, &itemCategory.Active); err != nil {
		return nil, err
	}
	if parentUID.Valid {
		uid, err := uuid.Parse(parentUID.String)
		if err != nil {
			return nil, err
		}
		itemCategory.ParentItemCategoryUID = &uid
	} else {
		itemCategory.ParentItemCategoryUID = nil
	}
	return &itemCategory, nil
}
