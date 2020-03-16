package data

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

var exJSON string = "https://jsonplaceholder.typicode.com/todos/1"

func FetchMain() {
	b := FetchAPI(exJSON)
	fmt.Println(b)
}

func FetchAPI(s string) Data {
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
	UserID    int
	ID        int
	Title     string
	Completed bool
}
