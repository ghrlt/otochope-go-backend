package helpers

import (
	"otochope/database"

	"github.com/google/uuid"
)

func GetAllCoupons() ([]Coupon, error) {
	query := "SELECT * FROM coupons"
	rows, err := database.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return parseCouponsRows(rows)
}

func GetActiveCoupons() ([]Coupon, error) {
	query := "SELECT * FROM coupons WHERE active = 1"
	rows, err := database.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return parseCouponsRows(rows)
}

func GetCouponByUID(uid uuid.UUID) (*Coupon, error) {
	query := "SELECT * FROM coupons WHERE uid = ? LIMIT 1"
	rows, err := database.Query(query, uid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return parseCouponRow(rows)
}

func GetCouponByCode(code string) (*Coupon, error) {
	query := "SELECT * FROM coupons WHERE code = ? LIMIT 1"
	rows, err := database.Query(query, code)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return parseCouponRow(rows)
}
