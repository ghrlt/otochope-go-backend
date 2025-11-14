package helpers

import "otochope/database"

func GetAllPlatforms() ([]Platform, error) {
	query := "SELECT * FROM platforms"
	rows, err := database.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return parsePlatformsRows(rows)
}

func GetPlatformByID(id int64) (*Platform, error) {
	query := "SELECT * FROM platforms WHERE id = ?"
	rows, err := database.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return parsePlatformRow(rows)
}

func GetPlatformByName(name string) (*Platform, error) {
	query := "SELECT * FROM platforms WHERE name = ?"
	rows, err := database.Query(query, name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return parsePlatformRow(rows)
}
