package models

import (
	"code.google.com/p/go-uuid/uuid"
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/lib/pq"
	"io"
	"regexp"
)

var DB *sql.DB

type Storage interface {
	WriteCloser(id string) (io.WriteCloser, error)
	ReadCloser(id string) (io.ReadCloser, error)
	Move(from, to string) error
	Delete(id string) error
	FileExists(id string) bool
}

var backend Storage

// Initialize the database
// steps:
// 1: create the user table
// 2: create the buckets table that depends on the user
// 3: create the objects table that depends on the buckets
// then create 1 admin user if we dont already have one
func Initialize(creds string, s Storage) error {
	if s == nil {
		return errors.New("i need a storage device")
	}
	backend = s
	// db, err := sql.Open("postgres", "dbname=test sslmode=disable")
	database, err := sql.Open("postgres", creds)
	if err != nil {
		return err
	}
	DB = database

	// create the user table
	_, err = DB.Exec("CREATE TABLE IF NOT EXISTS users (id uuid PRIMARY KEY,key character(10) NOT NULL, maxsize bigint DEFAULT 0, admin bool DEFAULT FALSE)")
	if err != nil {
		return err
	}

	// create the buckets table
	_, err = DB.Exec("CREATE TABLE IF NOT EXISTS buckets (id uuid PRIMARY KEY, name character varying(100) NOT NULL, user_id uuid references users(id) NOT NULL, UNIQUE (user_id, name))")
	if err != nil {
		return err
	}

	// create the objects table
	_, err = DB.Exec("CREATE TABLE IF NOT EXISTS objects (id uuid PRIMARY KEY, alias character varying(255) NOT NULL, size bigint, checksum character(32), bucket_id uuid references buckets(id) NOT NULL, public boolean DEFAULT FALSE, UNIQUE (bucket_id, alias))")
	if err != nil {
		return err
	}

	// create a admin user if we dont have one already
	rows, err := DB.Query("SELECT id, key, maxsize, admin FROM users WHERE admin = true")
	if err != nil {
		return err
	}
	defer rows.Close()
	count := 0
	for rows.Next() {
		u := User{}
		err = rows.Scan(&u.ID, &u.Key, &u.Limit, &u.Admin)
		if err == nil {
			count += 1
			fmt.Printf("admin user: %+v\n", u)
		}
	}
	if count == 0 {
		usr := User{
			ID:    generateID(),
			Key:   generateKey(),
			Limit: 0,
			Admin: true,
		}
		stmt, err := DB.Prepare("INSERT INTO users (id, key, maxsize, admin) VALUES ($1, $2, $3, $4)")
		if err == nil {
			stmt.Exec(usr.ID, usr.Key, usr.Limit, usr.Admin)
			fmt.Printf("creating admin user:%+v\n", usr)
		}
	}

	return nil
}

func StartCleaner() {

}

func uid(val string) string {
	if isId(val) {
		return val
	}
	return generateID()
}

func isId(val string) bool {
	matched, err := regexp.MatchString(`^\w\w\w\w\w\w\w\w-\w\w\w\w-\w\w\w\w-\w\w\w\w-\w\w\w\w\w\w\w\w\w\w\w\w$`, val)
	if err != nil {
		return false
	}
	return matched
}

func generateID() string {
	return uuid.New()
}
