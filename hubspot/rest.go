package main

import (
	"encoding/json"
	"fmt"
	"github.com/kr/pretty"
	"io/ioutil"
	"net/http"
)

func main() {
	//fmt.Println("Hello, the time is ",time.Now())
	getCall()
}


type Location struct {
	Name string
}

func getCall() {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "https://api.weatherapi.com/v1/current.json", nil)
	//req.Header.Add("Accept", "application/json")
	req.Header.Add("key", "b589960a9f184f53b4b224309212409")
	w := req.URL.Query()
	w.Add("q","89012")
	req.URL.RawQuery = w.Encode()
	fmt.Println(req.URL.String())
	response, err := client.Do(req)

	if err != nil {
		fmt.Printf("Failed.. Error: %s\n", err)
	} else {
		//defer response.Body.Close()
		data, _ := ioutil.ReadAll(response.Body)
		jsondata := string(data)
		fmt.Println(jsondata)
		//works until this point

		var rep Location

		json.Unmarshal(byte(jsondata['location']), &rep)
		pretty.Println(rep)

	}
}
