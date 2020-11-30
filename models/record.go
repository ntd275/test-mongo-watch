package models

import (
	"time"
)

type Record struct {
	Id           string      `bson:"_id"`
	Data         interface{} `bson:"data"`
	LastModified time.Time   `bson:"lastmodified"`
}
