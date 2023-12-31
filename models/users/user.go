package users

import (
	"database/sql"
	"time"

	"bench/logger"
	"bench/postgres"

	"github.com/google/uuid"
)

var pg = postgres.New()
var log = logger.New("bench:models:users")

type User struct {
	ID        *string    `json:"id"`
	Type      Type       `json:"type"`
	Name      *string    `json:"name"`
	Email     *string    `json:"email"`
	Password  *string    `json:"-"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

func New(name string, email string, password string) *User {
	now := time.Now()
	self := User{
		Type:      USER,
		Name:      &name,
		Email:     &email,
		Password:  &password,
		CreatedAt: &now,
		UpdatedAt: &now,
	}

	return &self
}

func GetByID(id string) *User {
	v := User{}
	err := pg.QueryRow(
		`
			SELECT *
			FROM users
			WHERE id = $1
		`,
		id,
	).Scan(
		&v.ID,
		&v.Type,
		&v.Name,
		&v.Email,
		&v.Password,
		&v.CreatedAt,
		&v.UpdatedAt,
		&v.DeletedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil
		}

		log.Error(err.Error())
	}

	return &v
}

func GetByEmail(email string) *User {
	v := User{}
	err := pg.QueryRow(
		`
			SELECT *
			FROM users
			WHERE email = $1
		`,
		email,
	).Scan(
		&v.ID,
		&v.Type,
		&v.Name,
		&v.Email,
		&v.Password,
		&v.CreatedAt,
		&v.UpdatedAt,
		&v.DeletedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil
		}

		log.Error(err.Error())
	}

	if v.DeletedAt != nil {
		return nil
	}

	return &v
}

func (self *User) Save() {
	if self.ID == nil {
		self.create()
	} else {
		self.update()
	}
}

func (self *User) Delete() {
	now := time.Now()
	self.DeletedAt = &now
	_, err := pg.Exec(
		`
			UPDATE users SET deleted_at = $2
			WHERE id = $1
		`,
		self.ID, now,
	)

	if err != nil {
		log.Error(err.Error())
	}
}

func (self *User) create() {
	id := uuid.New().String()
	self.ID = &id
	_, err := pg.Exec(
		`
			INSERT INTO users (id, type, name, email, password, created_at, updated_at)
			VALUES ($1, $2, $3, $4, $5, $6, $7)
		`,
		self.ID,
		self.Type,
		self.Name,
		self.Email,
		self.Password,
		self.CreatedAt,
		self.UpdatedAt,
	)

	if err != nil {
		log.Error(err.Error())
	}
}

func (self *User) update() {
	now := time.Now()
	self.UpdatedAt = &now
	_, err := pg.Exec(
		`
			UPDATE users SET type = $2, name = $3, email = $4, password = $5, updated_at = $6
			WHERE id = $1
		`,
		self.ID,
		self.Type,
		self.Name,
		self.Email,
		self.Password,
		now,
	)

	if err != nil {
		log.Error(err.Error())
	}
}
