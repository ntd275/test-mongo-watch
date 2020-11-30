package common

import (
	"errors"
	"os"
)

var (
	ErrorNotFound = errors.New("Not Found")
	ErrorStale    = errors.New("Not Updated Record")
)

func GetEnv(key string, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
