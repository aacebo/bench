package languages

import (
	"bench/logger"
	"bench/postgres"
	"time"
)

var pg = postgres.New()
var log = logger.New("bench:models:languages")

type Language struct {
	ID           *string    `json:"id"`
	IconURL      *string    `json:"icon_url"`
	Name         *string    `json:"name"`
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
