package main

//NJSON is a cool library to learn
//https://dev.to/m7shapan/golang-how-to-unmarshal-a-subset-of-nested-json-data-d84


import (
  "encoding/json"
  "fmt"
	"log"
)


type Location struct {
	Name string

}

type Current struct  {
  Temp_F float32
}

type Info struct {
  Location Location
  Current Current
}


func getpt(jsondata string){
  var dt = []byte(jsondata)

  var info Info
    if err := json.Unmarshal(dt, &info); err != nil {
      log.Fatal(err)
    }
    // fmt.Printf("%+v\n", info)
    fmt.Println(info.Location.Name)
    fmt.Println(info.Current.Temp_F)

}
