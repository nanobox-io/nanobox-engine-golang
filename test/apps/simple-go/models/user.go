package models

import (
	"fmt"
)

type User struct {
	ID    string
	Key   string
	Limit int64
	Admin bool
}

func (self *User) LimitExceeded() bool {
	if self.Limit == 0 {
		return false
	}

	rows, err := DB.Query("SELECT COALESCE(sum(objects.size), 0) FROM users JOIN buckets ON (buckets.user_id = users.id) LEFT OUTER JOIN objects ON (objects.bucket_id = buckets.id) WHERE users.id = $1 GROUP BY users.id", self.ID)
	if err != nil {
		return false
	}
	defer rows.Close()

	var size int64

	for rows.Next() {
		err = rows.Scan(&size)
		if err != nil {
			return false
		}
		break
	}
	return (self.Limit < size)
}

func GetUser(id string) (*User, error) {
	rows, err := DB.Query("SELECT id, key, maxsize, admin FROM users WHERE id = $1", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	u := User{}
	for rows.Next() {
		err = rows.Scan(&u.ID, &u.Key, &u.Limit, &u.Admin)
		if err != nil {
			return nil, err
		}
	}
	return &u, nil
}

func CreateUser(size int64) (*User, error) {
	usr := User{
		ID:    generateID(),
		Key:   generateKey(),
		Limit: size,
		Admin: false,
	}

	stmt, err := DB.Prepare("INSERT INTO users (id, key, maxsize, admin) VALUES ($1, $2, $3, FALSE)")
	if err != nil {
		return nil, err
	}
	_, err = stmt.Exec(usr.ID, usr.Key, usr.Limit)
	if err != nil {
		return nil, err
	}
	return &usr, nil
}

func DeleteUser(id string) error {
	// Delete
	stmt, err := DB.Prepare("delete from users where id=$1")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(id)
	if err != nil {
		return err
		fmt.Println(err)
	}
	return nil
}

func ListUsers() (*[]User, error) {
	users := []User{}
	rows, err := DB.Query("SELECT id, key, maxsize, admin FROM users")
	if err != nil {
		return &users, err
	}
	defer rows.Close()

	for rows.Next() {
		u := User{}
		err := rows.Scan(&u.ID, &u.Key, &u.Limit, &u.Admin)
		if err != nil {
			return &users, err
		}
		users = append(users, u)
	}
	return &users, nil
}

func generateKey() string {
	return generateID()[26:]
}
