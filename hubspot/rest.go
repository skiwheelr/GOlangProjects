package main

import (
	"fmt"
	"github.com/kr/pretty"
	"io/ioutil"
	"net/http"
)

func main() {
	//fmt.Println("Hello, the time is ",time.Now())
	retval := getCall()
	// fmt.Println(retval)
	// getpt(retval)
	// getptmap(retval)
	temppair := getptgj(retval)
	// fmt.Println(temppair)
	cval := multiplytemp(temppair)
	fmt.Println("Temp in Centigrade is", cval)


}

func getCall() string {
	jsondata := string("n/a")
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
		jsondata = string(data)
		pretty.Println("Json Received")
		// pretty.Println(jsondata)
	}
	return jsondata
}
