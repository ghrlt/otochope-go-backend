package helpers

import (
	"otochope/database"

	"github.com/google/uuid"
)

type CouponStats struct {
	TotalCoupons    int
	ActiveCoupons   int
	InactiveCoupons int
}

func GetCouponsStatistics() (*CouponStats, error) {
	query := `
		SELECT
			COUNT(*) AS total,
			SUM(CASE WHEN active = 1 THEN 1 ELSE 0 END) AS active,
			SUM(CASE WHEN active = 0 THEN 1 ELSE 0 END) AS inactive
		FROM coupons;
	`

	rows, err := database.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var stats CouponStats
	if rows.Next() {
		if err := rows.Scan(&stats.TotalCoupons, &stats.ActiveCoupons, &stats.InactiveCoupons); err != nil {
			return nil, err
		}
	}

	return &stats, nil
}

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
