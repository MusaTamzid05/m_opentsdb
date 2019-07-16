package main

import (
	"fmt"
	"log"
	"opentsdb/db"
	"time"
)

func addDataExample() {

	database, err := db.NewOpenTSDB("0.0.0.0:4242")

	if err != nil {
		log.Fatalln(err)
	}

	tags := map[string]string{
		"host": "Localhost",
	}

	for i := 0; i < 10; i++ {

		time.Sleep(1 * time.Second)
		data := make(map[string]interface{})
		data["Temp"] = i
		database.Insert(tags, data)
	}

}

func searchExample() {

	database, err := db.NewOpenTSDB("0.0.0.0:4242")

	if err != nil {
		log.Fatalln(err)
	}

	tags := map[string]string{
		"host": "Localhost",
	}

	start := time.Date(
		2019, 7, 16, 20, 0, 58, 651387237, time.UTC).Unix()

	end := time.Now().Unix()

	queryParams := db.MakeQuery(start, end, "Temp", tags)

	results, err := database.Search(queryParams)

	if err != nil {
		log.Fatalln(err)
	}

	for metrixName, result := range results {

		fmt.Println("*****************************")
		fmt.Println("Showing all the of Metrix :", metrixName)

		for i, data := range result {
			log.Println(i+1, ":", data)
		}

		fmt.Println("*****************************")
	}

}
