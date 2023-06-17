package problems

import (
	"bench/logger"
	"bench/postgres"
	"database/sql"
	"time"

	"github.com/google/uuid"
)

var pg = postgres.New()
var log = logger.New("bench:models:problems")

type Problem struct {
	ID          *string    `json:"id"`
	Name        *string    `json:"name"`
	DisplayName *string    `json:"display_name"`
	Description *string    `json:"description"`
	CreatedById *string    `json:"created_by_id"`
	CreatedAt   *time.Time `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
}

func New(name string, displayName string, description string, createdById string) *Problem {
	now := time.Now()
	self := Problem{
		Name:        &name,
		DisplayName: &displayName,
		Description: &description,
		CreatedById: &createdById,
		CreatedAt:   &now,
		UpdatedAt:   &now,
	}

	return &self
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
			&v.DisplayName,
			&v.Description,
			&v.CreatedById,
			&v.CreatedAt,
			&v.UpdatedAt,
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
		&v.DisplayName,
		&v.Description,
		&v.CreatedById,
		&v.CreatedAt,
		&v.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil
		}

		log.Error(err.Error())
	}

	return &v
}

func (self *Problem) Save() {
	if self.ID == nil {
		self.create()
	} else {
		self.update()
	}
}

func (self *Problem) create() {
	id := uuid.New().String()
	self.ID = &id
	_, err := pg.Exec(
		`
			INSERT INTO problems (id, name, display_name, description, created_by_id, created_at, updated_at)
			VALUES ($1, $2, $3, $4, $5, $6, $7)
		`,
		self.ID,
		self.Name,
		self.DisplayName,
		self.Description,
		self.CreatedById,
		self.CreatedAt,
		self.UpdatedAt,
	)

	if err != nil {
		log.Error(err.Error())
	}
}

func (self *Problem) update() {
	now := time.Now()
	self.UpdatedAt = &now
	_, err := pg.Exec(
		`
			UPDATE problems SET name = $2, display_name = $3, description = $4, updated_at = $5
			WHERE id = $1
		`,
		self.ID,
		self.Name,
		self.DisplayName,
		self.Description,
		now,
	)

	if err != nil {
		log.Error(err.Error())
	}
}
