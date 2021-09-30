package main

import(
  "fmt"
  "github.com/tidwall/gjson"

)

const jsonx = `{"name":{"first":"Janet","last":"Prichard"},"age":47}`

func getexample() {
  	value := gjson.Get(jsonx, "name.last")
  	fmt.Println(value.String())
}

func getptgj(jsondata string) map[string]float64 {
  namevalue := "city" //gjson.Get(jsondata, "location.name")
  tempvalue := gjson.Get(jsondata, "current.temp_f")
  var wepair = make(map[string]float64)
  wepair[namevalue] = tempvalue.Float()
  return wepair
}

func multiplytemp(wepair2 map[string]float64) string {
  ftemp := wepair2["city"]
  fmt.Println("FTemp in City is", ftemp)
  ctemp := (ftemp-30.0)/2.0
  ctempString := fmt.Sprintf("%f", ctemp)
  return ctempString
}



// func multiplytemp(wepair2 map[string]float64) float64 {
// //safety here
// 	if name, ok := wepair2["Henderson"]; ok {
// 		fmt.Printf("name = %s, ok = %v\n", name, ok)
//     fmt.Println(wepair2["Henderson"])
//     return wepair2["Henderson"]
// 	} else {
// 		fmt.Printf("City %d not found\n", "Henderson")
//     fmt.Println("8.4")
//     return 8.4
// 	}
//
// }
