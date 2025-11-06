package helpers

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Coupon struct {
	UID               uuid.UUID
	Code              string
	Value             int64
	IsValueAPercent   bool
	TotalUsageLimit   *int64
	PerUserUsageLimit *int64
	MinCartTotal      *int64
	ValidFrom         sql.NullTime
	ValidUntil        sql.NullTime
	CreatedAt         time.Time
	Active            bool
}

func parseCouponsRows(rows *sql.Rows) ([]Coupon, error) {
	var coupons []Coupon
	for rows.Next() {
		var coupon Coupon
		if err := rows.Scan(
			&coupon.UID,
			&coupon.Code,
			&coupon.Value,
			&coupon.IsValueAPercent,
			&coupon.TotalUsageLimit,
			&coupon.PerUserUsageLimit,
			&coupon.MinCartTotal,
			&coupon.ValidFrom,
			&coupon.ValidUntil,
			&coupon.CreatedAt,
			&coupon.Active,
		); err != nil {
			return nil, err
		}
		coupons = append(coupons, coupon)
	}
	return coupons, nil
}

func parseCouponRow(rows *sql.Rows) (*Coupon, error) {
	var coupon Coupon
	if !rows.Next() {
		return nil, nil // No coupon found
	}
	if err := rows.Scan(
		&coupon.UID,
		&coupon.Code,
		&coupon.Value,
		&coupon.IsValueAPercent,
		&coupon.TotalUsageLimit,
		&coupon.PerUserUsageLimit,
		&coupon.MinCartTotal,
		&coupon.ValidFrom,
		&coupon.ValidUntil,
		&coupon.CreatedAt,
		&coupon.Active,
	); err != nil {
		return nil, err
	}
	return &coupon, nil
}
