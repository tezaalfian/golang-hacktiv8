package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	for {
		data := map[string]interface{}{
			"water": rand.Intn(100) + 1,
			"wind":  rand.Intn(100) + 1,
		}

		requestJson, err := json.Marshal(data)
		client := &http.Client{}
		if err != nil {
			log.Fatalln(err)
		}

		req, err := http.NewRequest("POST", "https://jsonplaceholder.typicode.com/posts", bytes.NewBuffer(requestJson))
		req.Header.Set("Content-Type", "application/json")
		if err != nil {
			log.Fatalln(err)
		}

		res, err := client.Do(req)
		if err != nil {
			log.Fatalln(err)
		}
		defer res.Body.Close()

		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Fatalln(err)
		}
		// ubah body menjadi interface dan hapus index id
		var bodyInterface map[string]interface{}
		json.Unmarshal(body, &bodyInterface)
		delete(bodyInterface, "id")
		var waterStatus string
		var windStatus string
		// cek jumlah water
		if bodyInterface["water"].(float64) > 8 {
			waterStatus = "Bahaya"
		} else if bodyInterface["water"].(float64) > 6 {
			waterStatus = "Siaga"
		} else {
			waterStatus = "Aman"
		}
		// cek jumlah wind
		if bodyInterface["wind"].(float64) > 15 {
			windStatus = "Bahaya"
		} else if bodyInterface["wind"].(float64) > 7 {
			windStatus = "Siaga"
		} else {
			windStatus = "Aman"
		}
		// ubah bodyInterface menjadi json
		body, err = json.Marshal(bodyInterface)
		if err != nil {
			log.Fatalln(err)
		}
		// print body dan status
		log.Println(string(body))
		log.Println("Status water : ", waterStatus)
		log.Println("Status wind : ", windStatus)
		time.Sleep(15 * time.Second)
	}
}
