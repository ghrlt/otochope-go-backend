package helpers

import (
	"database/sql"
)

type Platform struct {
	ID      int64
	Name    string
	LogoURL string
}

func parsePlatformsRows(rows *sql.Rows) ([]Platform, error) {
	var platforms []Platform
	for rows.Next() {
		var platform Platform
		if err := rows.Scan(&platform.ID, &platform.Name, &platform.LogoURL); err != nil {
			return nil, err
		}
		platforms = append(platforms, platform)
	}
	return platforms, nil
}

func parsePlatformRow(rows *sql.Rows) (*Platform, error) {
	var platform Platform
	if !rows.Next() {
		return nil, nil // No platform found
	}
	if err := rows.Scan(&platform.ID, &platform.Name, &platform.LogoURL); err != nil {
		return nil, err
	}
	return &platform, nil
}
