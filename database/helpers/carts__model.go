package helpers

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Cart struct {
	UID       uuid.UUID
	UserUID   uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	Active    bool
}

func (c *Cart) GetUser() (*User, error) {
	return GetUserByUID(c.UserUID)
}

func (c *Cart) GetLinkedCoupons() ([]CartCoupon, error) {
	return GetAllCartCouponsByCartUID(c.UID)
}

func (c *Cart) GetActiveLinkedCoupons() ([]CartCoupon, error) {
	return GetActiveCartCouponsByCartUID(c.UID)
}

func (c *Cart) GetItems() ([]CartItem, error) {
	return GetAllCartItemsByCartUID(c.UID)
}

func (c *Cart) GetActiveItems() ([]CartItem, error) {
	return GetActiveCartItemsByCartUID(c.UID)
}

func (c *Cart) GetLinkedTransactions() ([]Transaction, error) {
	return GetAllTransactionsByCartUID(c.UID)
}

func parseCartRows(rows *sql.Rows) ([]Cart, error) {
	var carts []Cart
	for rows.Next() {
		var cart Cart
		if err := rows.Scan(&cart.UID, &cart.UserUID, &cart.CreatedAt, &cart.UpdatedAt, &cart.Active); err != nil {
			return nil, err
		}
		carts = append(carts, cart)
	}
	return carts, nil
}

func parseCartRow(rows *sql.Rows) (*Cart, error) {
	var cart Cart
	if !rows.Next() {
		return nil, nil // No cart found
	}
	if err := rows.Scan(&cart.UID, &cart.UserUID, &cart.CreatedAt, &cart.UpdatedAt, &cart.Active); err != nil {
		return nil, err
	}
	return &cart, nil
}
