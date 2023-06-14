package postgres

import (
	"bench/logger"
	"bench/utils"
	"database/sql"
	"time"

	_ "github.com/lib/pq"
)

var log = logger.New("bench:postgres")
var conn *sql.DB

func New() *sql.DB {
	if conn == nil {
		url := utils.GetEnv(
			"POSTGRES_CONNECTION_STRING",
			"postgresql://admin:admin@localhost:5432/bench?sslmode=disable",
		)

		db, err := sql.Open("postgres", url)

		if err != nil {
			log.Error(err.Error())
		}

		db.SetMaxOpenConns(25)
		db.SetMaxIdleConns(25)
		db.SetConnMaxLifetime(5 * time.Minute)

		if err := db.Ping(); err != nil {
			log.Error(err.Error())
		}

		conn = db
		log.Infoln("connection established...")
	}

	return conn
}
