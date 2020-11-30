package main

import (
	"client/common"
	"client/models"
	"log"
	"net/http"
	"time"
)

var (
	dbServer string
)

func init() {
	dbServer = common.GetEnv("DB_SERVER", "http://localhost:8080/api/v1")
}

func GetRecord(id string) (res models.Record, err error) {
	client := http.Client{
		Timeout: 10 * time.Second,
	}

	request, err := http.NewRequest("GET", dbServer+"/"+id, nil)
	if err != nil {
		return
	}
	resp, err := client.Do(request)
	if err != nil {
		return
	}
	if resp.StatusCode != 200 {
		err = common.ErrorNotFound
	}
	return
}

func main() {
	start := time.Now()
	for i := 1; i <= 1000; i++ {
		GetRecord("1")
	}
	log.Println(time.Since(start))
}
