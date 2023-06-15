package sessions

import (
	"bench/logger"
	"bench/postgres"
	"database/sql"
	"time"

	"github.com/google/uuid"
)

var pg = postgres.New()
var log = logger.New("bench:models:sessions")

type Session struct {
	ID        *string    `json:"id"`
	UserID    *string    `json:"user_id"`
	CreatedAt *time.Time `json:"created_at"`
}

func New(userId string) *Session {
	now := time.Now()
	self := Session{
		UserID:    &userId,
		CreatedAt: &now,
	}

	return &self
}

func Ping() bool {
	_, err := pg.Exec(`SELECT 1 FROM sessions limit 1`)

	if err != nil {
		return false
	}

	return true
}

func GetByID(id string) *Session {
	v := Session{}
	err := pg.QueryRow(
		`
			SELECT *
			FROM sessions
			WHERE id = $1
		`,
		id,
	).Scan(
		&v.ID,
		&v.UserID,
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

func GetByUserID(userId string) *Session {
	v := Session{}
	err := pg.QueryRow(
		`
			SELECT *
			FROM sessions
			WHERE user_id = $1
		`,
		userId,
	).Scan(
		&v.ID,
		&v.UserID,
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

func (self *Session) Save() {
	if self.ID == nil {
		self.create()
	}
}

func (self *Session) Delete() {
	_, err := pg.Exec(
		`
			DELETE FROM sessions
			WHERE id = $1
		`,
		self.ID,
	)

	if err != nil {
		log.Error(err.Error())
	}
}

func (self *Session) create() {
	id := uuid.New().String()
	self.ID = &id
	_, err := pg.Exec(
		`
			INSERT INTO sessions (id, user_id, created_at)
			VALUES ($1, $2, $3)
		`,
		self.ID,
		self.UserID,
		self.CreatedAt,
	)

	if err != nil {
		log.Error(err.Error())
	}
}
