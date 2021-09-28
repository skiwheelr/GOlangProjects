package main

import (
    "encoding/json"
    "fmt"
    // "reflect"
)




func getptmap(jsondata string) {

  // Create a map to parse the JSON
	var data map[string]interface{}
  // var subdata map[string]interface{}

	// Parse our JSON string
	err := json.Unmarshal([]byte(jsondata), &data)
	if err != nil {
		fmt.Println("Error parsing JSON string - %s", err)
	}

  //having issues going beyond this point

  // subd := fmt.Sprintf("%v", data["location"])
  // fmt.Println(subd)
  //
  // err2 := json.Unmarshal([]byte(subd), &subdata)
	// if err2 != nil {
	// 	fmt.Println("Error parsing JSON string - %s", err2)
	// }
  //
	// // Print out one of our JSON values
	// fmt.Printf("Name is %s", subdata["name"])

}
