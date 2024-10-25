package db_user_profiles

import (
	"crypto/tls"
	"os"
	"time"

	"github.com/go-pg/pg/v10"
)

type uid uint64
type hash [32]byte

type User struct {
	login        string
	passwordHash hash
	userId       uid
}

var (
	pgdb *pg.DB
)

func getTLSConfig() *tls.Config {
	pgSSLMode := os.Getenv("PGSSLMODE")
	if pgSSLMode == "disable" {
		return nil
	}
	return &tls.Config{
		InsecureSkipVerify: true,
	}
}

func pgOptions() *pg.Options {
	return &pg.Options{
		TLSConfig: getTLSConfig(),
		User:      os.Getenv("POSTGRES_LOGIN"),
		Password:  os.Getenv("POSTGRES_PASSWORD"),

		MaxRetries:      1,
		MinRetryBackoff: -1,

		DialTimeout:  30 * time.Second,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,

		PoolSize:           10,
		MaxConnAge:         10 * time.Second,
		PoolTimeout:        30 * time.Second,
		IdleTimeout:        10 * time.Second,
		IdleCheckFrequency: 100 * time.Millisecond,
	}
}

func InitializeUserDb() error {
	pgdb = pg.Connect(pgOptions())
	_, err := pgdb.Exec("CREATE SCHEMA IF NOT EXISTS db")
	if err != nil {
		return err
	}
	_, err = pgdb.Exec(`CREATE TABLE IF NOT EXISTS db.users(id serial, login varchar(500), passwordHash varchar(32))`)
	if err != nil {
		pgdb.Exec("DROP SCHEMA db")
	}
	return err
}

func DropUserDb() {
	pgdb.Exec("DROP TABLE db.users")
	pgdb.Exec("DROP SCHEMA db")
	pgdb.Close()
}
