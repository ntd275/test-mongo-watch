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
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		err = common.ErrorNotFound
	}
	return
}

func main() {
	start := time.Now()
	var countPass int
	var countFail int
	for i := 1; i <= 1000; i++ {

		if _, err := GetRecord("1"); err != nil {
			countFail++
		} else {
			countPass++
		}
	}
	t := time.Since(start)
	log.Println("Success: ", countPass)
	log.Println("Fail: ", countFail)
	log.Println(t)
	log.Println("TPS: ", float64(1000)*float64(time.Second)/float64(t))
}
