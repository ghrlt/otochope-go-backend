package helpers

import (
	"otochope/database"

	"github.com/google/uuid"
)

func GetAllUserInventories() ([]UserInventory, error) {
	query := "SELECT * FROM user_inventory"
	rows, err := database.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return parseUserInventoryRows(rows)
}

func GetActiveUserInventories() ([]UserInventory, error) {
	query := "SELECT * FROM user_inventory WHERE active = 1"
	rows, err := database.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return parseUserInventoryRows(rows)
}

func GetAllUserInventoriesByUserUID(userUID uuid.UUID) ([]UserInventory, error) {
	query := "SELECT * FROM user_inventory WHERE user_uid = ?"
	rows, err := database.Query(query, userUID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return parseUserInventoryRows(rows)
}

func GetActiveUserInventoryByUserUID(userUID uuid.UUID) (*UserInventory, error) {
	query := "SELECT * FROM user_inventory WHERE user_uid = ? AND active = 1 LIMIT 1"
	rows, err := database.Query(query, userUID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return parseUserInventoryRow(rows)
}

func GetUserInventoryByUID(uid uuid.UUID) (*UserInventory, error) {
	query := "SELECT * FROM user_inventory WHERE uid = ?"
	rows, err := database.Query(query, uid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return parseUserInventoryRow(rows)
}

func GetActiveUserInventoryByUID(uid uuid.UUID) (*UserInventory, error) {
	query := "SELECT * FROM user_inventory WHERE uid = ? AND active = 1"
	rows, err := database.Query(query, uid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return parseUserInventoryRow(rows)
}
