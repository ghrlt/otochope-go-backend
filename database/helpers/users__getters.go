package helpers

import (
	"otochope/database"

	"github.com/google/uuid"
)

func GetActiveUsers() ([]User, error) {
	query := "SELECT * FROM users WHERE active = 1"
	rows, err := database.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return parseUsersRows(rows)
}

func GetAllUsers() ([]User, error) {
	query := "SELECT * FROM users"
	rows, err := database.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return parseUsersRows(rows)
}

func GetUserByUID(uid uuid.UUID) (*User, error) {
	query := "SELECT * FROM users WHERE uid = ? LIMIT 1"
	rows, err := database.Query(query, uid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return parseUserRow(rows)
}

func GetActiveUserByUID(uid uuid.UUID) (*User, error) {
	query := "SELECT * FROM users WHERE uid = ? AND active = 1 LIMIT 1"
	rows, err := database.Query(query, uid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return parseUserRow(rows)
}

func GetUserByTelegramID(telegramID int64) (*User, error) {
	query := "SELECT * FROM users WHERE telegram_id = ? LIMIT 1"
	rows, err := database.Query(query, telegramID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return parseUserRow(rows)
}

func GetActiveUserByTelegramID(telegramID int64) (*User, error) {
	query := "SELECT * FROM users WHERE telegram_id = ? AND active = 1 LIMIT 1"
	rows, err := database.Query(query, telegramID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return parseUserRow(rows)
}
