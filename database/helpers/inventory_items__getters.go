package helpers

import (
	"otochope/database"

	"github.com/google/uuid"
)

func GetAllUserInventoryItems() ([]UserInventoryItem, error) {
	query := "SELECT * FROM user_inventory_items"
	rows, err := database.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return parseUserInventoryItemRows(rows)
}

func GetActiveUserInventoryItems() ([]UserInventoryItem, error) {
	query := "SELECT * FROM user_inventory_items WHERE active = 1"
	rows, err := database.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return parseUserInventoryItemRows(rows)
}

func GetAllUserInventoryItemsByInventoryUID(inventoryUID uuid.UUID) ([]UserInventoryItem, error) {
	query := "SELECT * FROM user_inventory_items WHERE inventory_uid = ?"
	rows, err := database.Query(query, inventoryUID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return parseUserInventoryItemRows(rows)
}

func GetActiveUserInventoryItemsByInventoryUID(inventoryUID uuid.UUID) ([]UserInventoryItem, error) {
	query := "SELECT * FROM user_inventory_items WHERE inventory_uid = ? AND active = 1"
	rows, err := database.Query(query, inventoryUID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return parseUserInventoryItemRows(rows)
}

func GetAllUserInventoryItemsByItemUID(itemUID uuid.UUID) ([]UserInventoryItem, error) {
	query := "SELECT * FROM user_inventory_items WHERE item_uid = ?"
	rows, err := database.Query(query, itemUID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return parseUserInventoryItemRows(rows)
}

func GetActiveUserInventoryItemsByItemUID(itemUID uuid.UUID) ([]UserInventoryItem, error) {
	query := "SELECT * FROM user_inventory_items WHERE item_uid = ? AND active = 1"
	rows, err := database.Query(query, itemUID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return parseUserInventoryItemRows(rows)
}

func GetUserInventoryItemByUID(uid uuid.UUID) (*UserInventoryItem, error) {
	query := "SELECT * FROM user_inventory_items WHERE uid = ? LIMIT 1"
	rows, err := database.Query(query, uid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return parseUserInventoryItemRow(rows)
}

func GetActiveUserInventoryItemByUID(uid uuid.UUID) (*UserInventoryItem, error) {
	query := "SELECT * FROM user_inventory_items WHERE uid = ? AND active = 1 LIMIT 1"
	rows, err := database.Query(query, uid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return parseUserInventoryItemRow(rows)
}
