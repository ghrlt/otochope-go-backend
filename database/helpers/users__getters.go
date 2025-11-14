package helpers

import (
	"otochope/database"

	"github.com/google/uuid"
)

type UserStats struct {
	TotalUsers    int
	ActiveUsers   int
	InactiveUsers int
}

func GetUsersStatistics() (*UserStats, error) {
	query := `
		SELECT
			COUNT(*) AS total,
			SUM(CASE WHEN active = 1 THEN 1 ELSE 0 END) AS active,
			SUM(CASE WHEN active = 0 THEN 1 ELSE 0 END) AS inactive
		FROM users;
	`

	rows, err := database.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var stats UserStats
	if rows.Next() {
		if err := rows.Scan(&stats.TotalUsers, &stats.ActiveUsers, &stats.InactiveUsers); err != nil {
			return nil, err
		}
	}

	return &stats, nil
}

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

func FindUsersByUsername(username string) ([]User, error) {
	query := "SELECT * FROM users WHERE username LIKE ?"
	rows, err := database.Query(query, username+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return parseUsersRows(rows)
}

func FindUsersByIdentifier(identifier string) ([]User, error) {
	query := "SELECT * FROM users WHERE identifier LIKE ?"
	rows, err := database.Query(query, identifier+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return parseUsersRows(rows)
}

func FindUsersByUsernameAndIdentifier(username, identifier string) ([]User, error) {
	query := "SELECT * FROM users WHERE username LIKE ? AND identifier LIKE ?"
	rows, err := database.Query(query, username+"%", identifier+"%")
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
