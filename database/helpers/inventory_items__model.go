package helpers

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type UserInventoryItem struct {
	UID              uuid.UUID
	UserInventoryUID uuid.UUID
	ItemUID          uuid.UUID
	Quantity         int
	CreatedAt        time.Time
	UpdatedAt        *time.Time
	Active           bool
}

func (ii *UserInventoryItem) GetInventory() (*UserInventory, error) {
	return GetUserInventoryByUID(ii.UserInventoryUID)
}

func (ii *UserInventoryItem) GetItem() (*Item, error) {
	return GetItemByUID(ii.ItemUID)
}

func parseUserInventoryItemRows(rows *sql.Rows) ([]UserInventoryItem, error) {
	var inventoryItems []UserInventoryItem
	for rows.Next() {
		var inventoryItem UserInventoryItem
		if err := rows.Scan(&inventoryItem.UID, &inventoryItem.UserInventoryUID, &inventoryItem.ItemUID, &inventoryItem.Quantity, &inventoryItem.CreatedAt, &inventoryItem.UpdatedAt, &inventoryItem.Active); err != nil {
			return nil, err
		}
		inventoryItems = append(inventoryItems, inventoryItem)
	}
	return inventoryItems, nil
}

func parseUserInventoryItemRow(rows *sql.Rows) (*UserInventoryItem, error) {
	var inventoryItem UserInventoryItem
	if !rows.Next() {
		return nil, nil // No inventory item found
	}
	if err := rows.Scan(&inventoryItem.UID, &inventoryItem.UserInventoryUID, &inventoryItem.ItemUID, &inventoryItem.Quantity, &inventoryItem.CreatedAt, &inventoryItem.UpdatedAt, &inventoryItem.Active); err != nil {
		return nil, err
	}
	return &inventoryItem, nil
}
