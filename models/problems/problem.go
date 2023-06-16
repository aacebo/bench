package problems

import (
	"bench/logger"
	"bench/postgres"
	"database/sql"
	"time"
)

var pg = postgres.New()
var log = logger.New("bench:models:problems")

type Problem struct {
	ID          *string    `json:"id"`
	Name        *string    `json:"name"`
	Description *string    `json:"description"`
	CreatedAt   *time.Time `json:"created_at"`
}

func Get() []*Problem {
	rows, err := pg.Query(
		`
			SELECT *
			FROM problems
		`,
	)

	if err != nil {
		log.Error(err.Error())
	}

	defer rows.Close()
	arr := []*Problem{}

	for rows.Next() {
		v := Problem{}
		err := rows.Scan(
			&v.ID,
			&v.Name,
			&v.Description,
			&v.CreatedAt,
		)

		if err != nil {
			log.Error(err.Error())
		}

		arr = append(arr, &v)
	}

	return arr
}

func GetByName(name string) *Problem {
	v := Problem{}
	err := pg.QueryRow(
		`
			SELECT *
			FROM problems
			WHERE name = $1
		`,
		name,
	).Scan(
		&v.ID,
		&v.Name,
		&v.Description,
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
