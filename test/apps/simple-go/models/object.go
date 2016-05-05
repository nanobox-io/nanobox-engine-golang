package models

import (
	"errors"
	"fmt"
	"io"
)

type Object struct {
	ID       string
	BucketID string
	Alias    string
	CheckSum string
	Public   bool
	Size     int64
}

func (self *Object) WriteCloser() (io.WriteCloser, error) {
	wc, err := backend.WriteCloser(self.ID)
	if err != nil {
		return nil, err
	}
	return wc, nil
}

func (self *Object) ReadCloser() (io.ReadCloser, error) {
	rc, err := backend.ReadCloser(self.ID)
	if err != nil {
		return nil, err
	}
	return rc, nil
}

func (self *Object) Move(newId string) error {
	return backend.Move(self.ID, newId)
}

func (self *Object) Remove() error {
	return backend.Delete(self.ID)
}

func (self *Object) Exists() bool {
	return backend.FileExists(self.ID)
}

func SaveObject(newObject *Object) error {
	stmt, err := DB.Prepare("UPDATE objects SET alias=$2, checksum=$3, size=$4, public=$5 WHERE id=$1")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(newObject.ID, newObject.Alias, newObject.CheckSum, newObject.Size, newObject.Public)

	stmt, err = DB.Prepare("INSERT INTO objects (id, alias, checksum, size, public, bucket_id) SELECT $1, $2, $3, $4, $5, $6 WHERE NOT EXISTS (SELECT 1 FROM objects WHERE id=$1)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(newObject.ID, newObject.Alias, newObject.CheckSum, newObject.Size, newObject.Public, newObject.BucketID)
	return nil
}

func GetPublicObject(userId, bucketId, id string) (*Object, error) {
	fmt.Println(userId, bucketId, id)
	rows, err := DB.Query("SELECT objects.id, objects.alias, objects.checksum, objects.size, objects.bucket_id, objects.public FROM users JOIN buckets ON (buckets.user_id = users.id) JOIN objects ON (objects.bucket_id = buckets.id) WHERE objects.public AND (objects.id = $1 OR objects.alias = $2) AND users.id = $3 AND (buckets.id = $4 OR buckets.name = $5)", uid(id), id, userId, uid(bucketId), bucketId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	f := Object{}
	for rows.Next() {
		err = rows.Scan(&f.ID, &f.Alias, &f.CheckSum, &f.Size, &f.BucketID, &f.Public)
		if err != nil {
			return nil, err
		}
		break
	}
	if f.ID == "" {
		return nil, errors.New("not found")
	}
	return &f, nil
}

func GetObject(userId, userKey, bucketId, id string) (*Object, error) {
	rows, err := DB.Query("SELECT objects.id, objects.alias, objects.checksum, objects.size, objects.bucket_id, objects.public FROM users JOIN buckets ON (buckets.user_id = users.id) JOIN objects ON (objects.bucket_id = buckets.id) WHERE (objects.id = $1 OR objects.alias = $2) AND users.id = $3 AND users.key = $4 AND (buckets.id = $5 OR buckets.name = $6)", uid(id), id, userId, userKey, uid(bucketId), bucketId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	f := Object{}
	for rows.Next() {
		err = rows.Scan(&f.ID, &f.Alias, &f.CheckSum, &f.Size, &f.BucketID, &f.Public)
		if err != nil {
			return nil, err
		}
		break
	}
	if f.ID == "" {
		return nil, errors.New("not found")
	}
	return &f, nil
}

func CreateObject(userId, userKey, bucketId, alias string) (*Object, error) {
	bucket, err := GetBucket(userId, userKey, bucketId)
	if err != nil {
		return nil, err
	}

	f := Object{
		ID:       generateID(),
		BucketID: bucket.ID,
		Alias:    alias,
	}

	if f.Alias == "" {
		f.Alias = f.ID
	}

	stmt, err := DB.Prepare("INSERT INTO objects (id, alias, checksum, size, bucket_id) VALUES ($1, $2, $3, $4, (SELECT buckets.id FROM buckets JOIN users ON (buckets.user_id = users.id) WHERE (buckets.id = $5 OR buckets.name = $6) AND users.id = $7 AND users.key = $8))")
	if err != nil {
		return nil, err
	}

	_, err = stmt.Exec(f.ID, f.Alias, f.CheckSum, f.Size, uid(f.BucketID), f.BucketID, userId, userKey)
	if err != nil {
		return nil, err
	}
	return &f, nil
}

func ListObjects(userId, userKey, bucketId string) (*[]Object, error) {
	fs := []Object{}
	rows, err := DB.Query("SELECT objects.id, objects.alias, objects.checksum, objects.size, objects.bucket_id, objects.public FROM users JOIN buckets ON (buckets.user_id = users.id) JOIN objects ON (objects.bucket_id = buckets.id) WHERE objects.size != 0 AND (buckets.id = $1 OR buckets.name = $2) AND users.id = $3 AND users.key = $4", uid(bucketId), bucketId, userId, userKey)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		f := Object{}
		err := rows.Scan(&f.ID, &f.Alias, &f.CheckSum, &f.Size, &f.BucketID, &f.Public)
		if err != nil {
			return nil, err
		}
		fs = append(fs, f)
	}
	return &fs, nil
}

func DeleteObject(userId, userKey, bucketId, id string) error {
	stmt, err := DB.Prepare("DELETE FROM objects WHERE id=$1 AND bucket_id = (SELECT buckets.id FROM buckets JOIN users ON (buckets.user_id = users.id) WHERE users.id = $3 AND users.key = $4 AND (buckets.id = $5 OR buckets.name = $6) AND (objects.id = $1 OR objects.alias = $2))")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(uid(id), id, userId, userKey, uid(bucketId), bucketId)
	if err != nil {
		return err
	}
	return nil
}

func CleanEmptyObjects() {
	rows, err := DB.Query("SELECT objects.id, objects.alias, objects.checksum, objects.size, objects.bucket_id, objects.public FROM objects WHERE size = 0")
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()

	for rows.Next() {
		o := Object{}
		err = rows.Scan(&o.ID, &o.Alias, &o.CheckSum, &o.Size, &o.BucketID, &o.Public)
		if err != nil {
			fmt.Println(err)
		}
		o.Remove()
		stmt, err := DB.Prepare("DELETE FROM objects WHERE id=$1")
		if err != nil {
			fmt.Println(err)
		}
		_, err = stmt.Exec(o.ID)
		if err != nil {
			fmt.Println(err)
		}
	}
}
