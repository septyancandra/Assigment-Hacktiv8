package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	//res, err := http.Get("https://jsonplaceholder.typicode.com/posts/1")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(res.Body)
	//body, err := ioutil.ReadAll(res.Body)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//defer res.Body.Close()
	//
	//sb := string(body)
	//log.Println(sb)
	data := map[string]interface{}{
		"title":  "Coba",
		"body":   "A1",
		"userId": 1,
	}
	requestJson, err := json.Marshal(data)
	client := &http.Client{}
	if err != nil {
		log.Fatal(err)
	}

	req, err := http.NewRequest("POST", "https://jsonplaceholder.typicode.com/posts", bytes.NewBuffer(requestJson))
	req.Header.Set("Content-type", "application/json")
	if err != nil {
		log.Fatal(err)
	}

	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(body))
}
