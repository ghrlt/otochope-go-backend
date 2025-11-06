package helpers

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type User struct {
	UID        uuid.UUID
	TelegramID int64
	CreatedAt  time.Time
	Active     bool
}

func (u *User) GetCarts() ([]Cart, error) {
	return GetAllCartsByUserUID(u.UID)
}

func (u *User) GetActiveCart() (*Cart, error) {
	return GetActiveCartByUserUID(u.UID)
}

func (u *User) GetInventories() ([]UserInventory, error) {
	return GetAllUserInventoriesByUserUID(u.UID)
}

func (u *User) GetActiveInventory() (*UserInventory, error) {
	return GetActiveUserInventoryByUserUID(u.UID)
}

func (u *User) GetTransactions() ([]Transaction, error) {
	return GetAllTransactionsByUserUID(u.UID)
}

func parseUsersRows(rows *sql.Rows) ([]User, error) {
	var users []User
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.UID, &user.TelegramID, &user.CreatedAt, &user.Active); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func parseUserRow(rows *sql.Rows) (*User, error) {
	var user User
	if !rows.Next() {
		return nil, nil // No user found
	}
	if err := rows.Scan(&user.UID, &user.TelegramID, &user.CreatedAt, &user.Active); err != nil {
		return nil, err
	}
	return &user, nil
}
