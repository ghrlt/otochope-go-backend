package helpers

import (
	"otochope/database"

	"github.com/google/uuid"
)

func GetAllCartCoupons() ([]CartCoupon, error) {
	query := "SELECT * FROM user_cart_coupons"
	rows, err := database.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return parseCartCouponsRows(rows)
}

func GetAllActiveCartCoupons() ([]CartCoupon, error) {
	query := "SELECT * FROM user_cart_coupons WHERE active = 1"
	rows, err := database.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return parseCartCouponsRows(rows)
}

func GetAllCartCouponsByCartUID(cartUID uuid.UUID) ([]CartCoupon, error) {
	query := "SELECT * FROM user_cart_coupons WHERE cart_uid = ?"
	rows, err := database.Query(query, cartUID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return parseCartCouponsRows(rows)
}

func GetActiveCartCouponsByCartUID(cartUID uuid.UUID) ([]CartCoupon, error) {
	query := "SELECT * FROM user_cart_coupons WHERE cart_uid = ? AND active = 1"
	rows, err := database.Query(query, cartUID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return parseCartCouponsRows(rows)
}

func GetAllCartCouponsByCouponUID(couponUID uuid.UUID) ([]CartCoupon, error) {
	query := "SELECT * FROM user_cart_coupons WHERE coupon_uid = ?"
	rows, err := database.Query(query, couponUID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return parseCartCouponsRows(rows)
}

func GetActiveCartCouponsByCouponUID(couponUID uuid.UUID) ([]CartCoupon, error) {
	query := "SELECT * FROM user_cart_coupons WHERE coupon_uid = ? AND active = 1"
	rows, err := database.Query(query, couponUID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return parseCartCouponsRows(rows)
}

func GetCartCouponByUID(uid uuid.UUID) (*CartCoupon, error) {
	query := "SELECT * FROM user_cart_coupons WHERE uid = ?"
	rows, err := database.Query(query, uid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return parseCartCouponRow(rows)
}

func GetActiveCartCouponByUID(uid uuid.UUID) (*CartCoupon, error) {
	query := "SELECT * FROM user_cart_coupons WHERE uid = ? AND active = 1"
	rows, err := database.Query(query, uid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return parseCartCouponRow(rows)
}
