package languages

import (
	"bench/logger"
	"bench/postgres"
	"database/sql"
	"time"
)

var pg = postgres.New()
var log = logger.New("bench:models:languages")

type Language struct {
	ID           *string    `json:"id"`
	IconURL      *string    `json:"icon_url"`
	Name         *string    `json:"name"`
	Version      *string    `json:"version"`
	IsConcurrent *bool      `json:"is_concurrent"`
	CreatedAt    *time.Time `json:"created_at"`
}

func Get() []*Language {
	rows, err := pg.Query(
		`
			SELECT *
			FROM languages
		`,
	)

	if err != nil {
		log.Error(err.Error())
	}

	defer rows.Close()
	arr := []*Language{}

	for rows.Next() {
		v := Language{}
		err := rows.Scan(
			&v.ID,
			&v.IconURL,
			&v.Name,
			&v.Version,
			&v.IsConcurrent,
			&v.CreatedAt,
		)

		if err != nil {
			log.Error(err.Error())
		}

		arr = append(arr, &v)
	}

	return arr
}

func GetByName(name string) *Language {
	v := Language{}
	err := pg.QueryRow(
		`
			SELECT *
			FROM languages
			WHERE name = $1
		`,
		name,
	).Scan(
		&v.ID,
		&v.IconURL,
		&v.Name,
		&v.Version,
		&v.IsConcurrent,
		&v.CreatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil
		}

		log.Error(err.Error())
	}

	return &v
}
