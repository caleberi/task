package lib

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/boltdb/bolt"
)

var db *bolt.DB
var bucket []byte

func InitDatabase(dbPath string, bucketName string) (*bolt.DB, error) {
	var err error
	bucket = []byte(bucketName)
	db, err = bolt.Open(dbPath, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		log.Fatal(err)
	}
	return db, db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(bucket)
		if err != nil {
			return fmt.Errorf("error [create bucket: %s]", err)
		}
		return nil
	})
}

func InsertInDatabase(task Task) (int, error) {
	var id int
	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucket)
		id64, _ := b.NextSequence()
		id = int(id64)
		task.Id = id
		val := fmt.Sprintf(
			"id=%d,description=%s,status=%s",
			task.Id,
			task.Description,
			task.GetStatusString())
		err := b.Put(Itob(id), []byte(val))
		return err
	})
	return id, err
}

func DeleteDataFromDatabase(taskId int) error {
	return db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucket)
		return b.Delete(Itob(taskId))
	})
}

func RetrieveFromDatabase(key int) (string, error) {
	var ret []byte
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucket)
		ret = b.Get(Itob(key))
		return nil
	})
	if err != nil {
		return "", err
	}
	return string(ret), nil
}

func ListDataFromDatabase() ([]Task, error) {
	var data []Task
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucket)
		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			val := string(v)
			tokens := strings.Split(val, ",")
			id := strings.Split(tokens[0], "=")[1]
			oid, err := strconv.Atoi(id)
			CheckError(err)
			description := strings.Split(tokens[1], "=")[1]
			statusStr := strings.Split(tokens[2], "=")[1]
			task := Task{
				Id:          oid,
				Description: description,
				Status:      GetStatusFromString(statusStr),
			}
			data = append(data, task)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return data, nil
}

func CloseDatabase(db *bolt.DB) {
	err := db.Close()
	CheckError(err)
}
