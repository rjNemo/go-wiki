package data

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

var exJson string = "https://jsonplaceholder.typicode.com/todos/1"

func Main() {
	b := fetchApi(exJson)
	fmt.Println(b)
}

func FetchApi(s string) Data {
	r, err := http.NewRequest(http.MethodGet, s, nil)
	if err != nil {
		log.Fatal(err)
	}
	Client := http.Client{
		Timeout: time.Second * 2, // Maximum of 2 secs
	}
	rs, err := Client.Do(r)
	b, err := ioutil.ReadAll(rs.Body)
	d := Data{}
	err = json.Unmarshal(b, &d)
	if err != nil {
		log.Fatal(err)
	}
	return d
}

type Data struct {
	UserId    int
	Id        int
	Title     string
	Completed bool
}
