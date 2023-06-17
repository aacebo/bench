package tests

import (
	"bench/logger"
	"bench/postgres"
	"database/sql"
	"time"

	"github.com/google/uuid"
)

var pg = postgres.New()
var log = logger.New("bench:models:tests")

type Test struct {
	ID        *string    `json:"id"`
	ProblemID *string    `json:"problem_id"`
	Input     *string    `json:"input"`
	Output    *string    `json:"output"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

func New(problemId string, input string, output string) *Test {
	now := time.Now()
	self := Test{
		ProblemID: &problemId,
		Input:     &input,
		Output:    &output,
		CreatedAt: &now,
		UpdatedAt: &now,
	}

	return &self
}

func GetByProblemID(problemId string) []*Test {
	rows, err := pg.Query(
		`
			SELECT *
			FROM tests
			WHERE problem_id = $1
		`,
		problemId,
	)

	if err != nil {
		log.Error(err.Error())
	}

	defer rows.Close()
	arr := []*Test{}

	for rows.Next() {
		v := Test{}
		err := rows.Scan(
			&v.ID,
			&v.ProblemID,
			&v.Input,
			&v.Output,
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

func GetByID(id string) *Test {
	v := Test{}
	err := pg.QueryRow(
		`
			SELECT *
			FROM tests
			WHERE id = $1
		`,
		id,
	).Scan(
		&v.ID,
		&v.ProblemID,
		&v.Input,
		&v.Output,
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

func (self *Test) Save() {
	if self.ID == nil {
		self.create()
	} else {
		self.update()
	}
}

func (self *Test) Delete() {
	_, err := pg.Exec(
		`
			DELETE FROM tests
			WHERE id = $1
		`,
		self.ID,
	)

	if err != nil {
		log.Error(err.Error())
	}
}

func (self *Test) create() {
	id := uuid.New().String()
	self.ID = &id
	_, err := pg.Exec(
		`
			INSERT INTO tests (id, problem_id, input, output, created_at, updated_at)
			VALUES ($1, $2, $3, $4, $5, $6)
		`,
		self.ID,
		self.ProblemID,
		self.Input,
		self.Output,
		self.CreatedAt,
		self.UpdatedAt,
	)

	if err != nil {
		log.Error(err.Error())
	}
}

func (self *Test) update() {
	now := time.Now()
	self.UpdatedAt = &now
	_, err := pg.Exec(
		`
			UPDATE tests SET input = $2, output = $3, updated_at = $4
			WHERE id = $1
		`,
		self.ID,
		self.Input,
		self.Output,
		now,
	)

	if err != nil {
		log.Error(err.Error())
	}
}
