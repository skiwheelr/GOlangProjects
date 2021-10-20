package main

import (
	"encoding/json"
	"fmt"
	"github.com/kr/pretty"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	fmt.Println("Hello, the time is ",time.Now())
	getCall()
}

type User struct {
	Name string
	// id int8
}

func getCall() {
	response, err := http.Get("http://jsonplaceholder.typicode.com/users")
	if err != nil {
		fmt.Printf("Failed.. Error: %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		jsondata := string(data)
		// fmt.Print(jsondata)
		var users []User
		json.Unmarshal([]byte(jsondata), &users)
		//fmt.Printf("Users: %+v", users)
		// users[0].Name = "Mark Wagner"
		// users[0].id = 100
		pretty.Println(users)
		// fmt.Print(users)
	}
}

/* Results:

Hello, the time is  2020-04-10 00:47:18.21 -0700 MST m=+0.000628811
[]main.User{
{Name:"Mark Wagner", id:100},
{Name:"Ervin Howell", id:0},
{Name:"Clementine Bauch", id:0},
{Name:"Patricia Lebsack", id:0},
{Name:"Chelsey Dietrich", id:0},
{Name:"Mrs. Dennis Schulist", id:0},
{Name:"Kurtis Weissnat", id:0},
{Name:"Nicholas Runolfsdottir V", id:0},
{Name:"Glenna Reichert", id:0},
{Name:"Clementina DuBuque", id:0},
}*/
