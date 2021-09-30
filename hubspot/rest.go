package main

import (
	"fmt"
	"github.com/kr/pretty"
	"io/ioutil"
	"net/http"
	"encoding/json"
	"bytes"
	"os"
)

func main() {
	arg := os.Args[1]
  // fmt.Println(arg)
	//fmt.Println("Hello, the time is ",time.Now())
	retval := getCall(arg)
	temppair := getptgj(retval)
	cval := multiplytemp(temppair)
	fmt.Println("Temp in Centigrade is", cval)
	postCall(cval)
}

func getCall(arg string) string {
	jsondata := string("n/a")
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "https://api.weatherapi.com/v1/current.json", nil)
	//req.Header.Add("Accept", "application/json")
	req.Header.Add("key", "b589960a9f184f53b4b224309212409")
	w := req.URL.Query()
	w.Add("q", arg)
	req.URL.RawQuery = w.Encode()
	fmt.Println(req.URL.String())
	response, err := client.Do(req)

	if err != nil {
		fmt.Printf("Failed.. Error: %s\n", err)
	} else {
		//defer response.Body.Close()
		data, _ := ioutil.ReadAll(response.Body)
		jsondata = string(data)
		pretty.Println("Json Received")
		pretty.Println(jsondata)
	}
	return jsondata
}

func postCall(result string) {

	postBody, _ := json.Marshal(map[string]string{
   "Centigrade_Temp": result,
	})
	responseBody := bytes.NewBuffer(postBody)

	client := &http.Client{}
	req, _ := http.NewRequest("POST", "https://httpbin.org/post", responseBody)
	req.Header.Add("Accept", "application/json")
	w := req.URL.Query()
	req.URL.RawQuery = w.Encode()
	fmt.Println(req.URL.String())
	response, err := client.Do(req)

	if err != nil {
		fmt.Printf("Failed.. Error: %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		jsonresdata := string(data)
		pretty.Println("Data Posted")
		pretty.Println(jsonresdata)
	}
}
