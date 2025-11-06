package helpers

import (
	"database/sql"

	"github.com/google/uuid"
)

type UserInventory struct {
	UID     uuid.UUID
	UserUID uuid.UUID
	Active  bool
}

func (ui *UserInventory) GetUser() (*User, error) {
	return GetUserByUID(ui.UserUID)
}

func (ui *UserInventory) GetItems() ([]UserInventoryItem, error) {
	return GetAllUserInventoryItemsByInventoryUID(ui.UID)
}

func (ui *UserInventory) GetActiveItems() ([]UserInventoryItem, error) {
	return GetActiveUserInventoryItemsByInventoryUID(ui.UID)
}

func parseUserInventoryRows(rows *sql.Rows) ([]UserInventory, error) {
	var userInventories []UserInventory
	for rows.Next() {
		var userInventory UserInventory
		if err := rows.Scan(&userInventory.UID, &userInventory.UserUID, &userInventory.Active); err != nil {
			return nil, err
		}
		userInventories = append(userInventories, userInventory)
	}
	return userInventories, nil
}

func parseUserInventoryRow(rows *sql.Rows) (*UserInventory, error) {
	var userInventory UserInventory
	if !rows.Next() {
		return nil, nil // No user inventory found
	}
	if err := rows.Scan(&userInventory.UID, &userInventory.UserUID, &userInventory.Active); err != nil {
		return nil, err
	}
	return &userInventory, nil
}
