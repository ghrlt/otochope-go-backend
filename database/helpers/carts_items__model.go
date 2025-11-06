package helpers

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type CartItem struct {
	UID       uuid.UUID
	CartUID   uuid.UUID
	ItemUID   uuid.UUID
	Quantity  int
	AddedAt   time.Time
	UpdatedAt time.Time
	Active    bool
}

func (ci *CartItem) GetCart() (*Cart, error) {
	return GetCartByUID(ci.CartUID)
}

func (ci *CartItem) GetItem() (*Item, error) {
	return GetItemByUID(ci.ItemUID)
}

func parseCartItemRows(rows *sql.Rows) ([]CartItem, error) {
	var cartItems []CartItem
	for rows.Next() {
		var cartItem CartItem
		if err := rows.Scan(&cartItem.UID, &cartItem.CartUID, &cartItem.ItemUID, &cartItem.Quantity, &cartItem.AddedAt, &cartItem.UpdatedAt, &cartItem.Active); err != nil {
			return nil, err
		}
		cartItems = append(cartItems, cartItem)
	}
	return cartItems, nil
}

func parseCartItemRow(rows *sql.Rows) (*CartItem, error) {
	var cartItem CartItem
	if !rows.Next() {
		return nil, nil // No cart item found
	}
	if err := rows.Scan(&cartItem.UID, &cartItem.CartUID, &cartItem.ItemUID, &cartItem.Quantity, &cartItem.AddedAt, &cartItem.UpdatedAt, &cartItem.Active); err != nil {
		return nil, err
	}
	return &cartItem, nil
}
