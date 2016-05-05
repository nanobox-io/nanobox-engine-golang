package models

import (
	"errors"
)

type SizedBucket struct {
	Bucket
	Size int64
}

type Bucket struct {
	ID     string
	UserID string
	Name   string
}

func GetBucket(userID, userKey, id string) (*SizedBucket, error) {
	rows, err := DB.Query("SELECT buckets.*, COALESCE(sum(objects.size), 0) FROM users JOIN buckets ON (buckets.user_id = users.id) LEFT OUTER JOIN objects ON (objects.bucket_id = buckets.id) WHERE (buckets.id = $1 OR buckets.name = $2) AND users.id = $3 AND users.key = $4 GROUP BY buckets.id", uid(id), id, userID, userKey)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	buck := SizedBucket{}
	for rows.Next() {
		err = rows.Scan(&buck.ID, &buck.Name, &buck.UserID, &buck.Size)
		if err != nil {
			return nil, err
		}
		break
	}
	if buck.ID == "" {
		return nil, errors.New("not found")
	}
	return &buck, err
}

func CreateBucket(userId, key, name string) (*Bucket, error) {
	buck := Bucket{
		ID:     generateID(),
		Name:   name,
		UserID: userId,
	}

	stmt, err := DB.Prepare("INSERT INTO buckets (id, name, user_id) VALUES ($1, $2, (SELECT id FROM users WHERE id = $3 and key = $4))")
	if err != nil {
		return nil, err
	}
	_, err = stmt.Exec(buck.ID, buck.Name, buck.UserID, key)
	if err != nil {
		return nil, err
	}
	return &buck, nil
}

func ListBuckets(userId, userKey string) (*[]Bucket, error) {
	bucks := []Bucket{}
	rows, err := DB.Query("SELECT * FROM buckets WHERE user_id = (SELECT id FROM users WHERE id = $1 and key = $2)", userId, userKey)
	if err != nil {
		return &bucks, err
	}
	defer rows.Close()

	for rows.Next() {
		b := Bucket{}
		err := rows.Scan(&b.ID, &b.Name, &b.UserID)
		if err != nil {
			return nil, err
		}
		bucks = append(bucks, b)
	}

	return &bucks, nil

}

func DeleteBucket(userId, userKey, id string) error {
	// Delete
	stmt, err := DB.Prepare("DELETE FROM buckets where (id=$1 OR name=$2) AND user_id=(SELECT id FROM users WHERE id = $3 and key = $4)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(uid(id), id, userId, userKey)
	if err != nil {
		return err
	}

	return nil
}
