package helpers

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type CartCoupon struct {
	UID       uuid.UUID
	CartUID   uuid.UUID
	CouponUID uuid.UUID
	AddedAt   time.Time
	Active    bool
}

func (cc *CartCoupon) GetCart() (*Cart, error) {
	return GetCartByUID(cc.CartUID)
}

func (cc *CartCoupon) GetCoupon() (*Coupon, error) {
	return GetCouponByUID(cc.CouponUID)
}

func parseCartCouponsRows(rows *sql.Rows) ([]CartCoupon, error) {
	var cartCoupons []CartCoupon
	for rows.Next() {
		var cartCoupon CartCoupon
		if err := rows.Scan(&cartCoupon.UID, &cartCoupon.CartUID, &cartCoupon.CouponUID, &cartCoupon.AddedAt, &cartCoupon.Active); err != nil {
			return nil, err
		}
		cartCoupons = append(cartCoupons, cartCoupon)
	}
	return cartCoupons, nil
}

func parseCartCouponRow(rows *sql.Rows) (*CartCoupon, error) {
	var cartCoupon CartCoupon
	if !rows.Next() {
		return nil, nil // No cart coupon found
	}
	if err := rows.Scan(&cartCoupon.UID, &cartCoupon.CartUID, &cartCoupon.CouponUID, &cartCoupon.AddedAt, &cartCoupon.Active); err != nil {
		return nil, err
	}
	return &cartCoupon, nil
}
