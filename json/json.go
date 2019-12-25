package main

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func main() {
	// json Marshal and Unmarshal bool
	boolS, _ := json.Marshal(true)
	fmt.Println(string(boolS))
	var b bool
	if err := json.Unmarshal(boolS, &b); err != nil {
		panic(err)
	}

	// json Marshal and Unmarshal float
	floatS, _ := json.Marshal(3.14159)
	fmt.Println(string(floatS))
	var f float64
	if err := json.Unmarshal(floatS, &f); err != nil {
		panic(err)
	}
	fmt.Println("unmarshaled float64:", f)

	// json Marshal nil
	nilS, _ := json.Marshal(nil)
	fmt.Println(string(nilS))

	// json Unmarshal nil
	var p interface{}
	if err := json.Unmarshal(nilS, &p); err != nil {
		panic(err)
	}
	fmt.Println("unmarshal null:", p)

	// json.Unmarshal loss of precision
	// json.Unmarshal use float64 for json numbers
	jsonStr := `{"userID":1,"config":{"pid":1234567890123456789,"target_type":1}}`
	var result map[string]interface{}
	json.Unmarshal([]byte(jsonStr), &result)
	fmt.Println("json.Unmarshal result is ", result)
	pid := result["config"].(map[string]interface{})["pid"]
	fmt.Printf("pid type is %T\n", pid)
	fmt.Printf("pid value is %f\n", pid)
	fmt.Printf("Int64 is %d\n", int64(pid.(float64)))

	// json.Decoder
	var resultEx map[string]interface{}
	decoder := json.NewDecoder(bytes.NewReader([]byte(jsonStr)))
	// UseNumber causes the Decoder to unmarshal a number into an interface{} as a Number instead of a float64
	decoder.UseNumber()
	decoder.Decode(&resultEx)
	fmt.Println("json.Unmarshal result is ", resultEx)
	pidEx := resultEx["config"].(map[string]interface{})["pid"]
	fmt.Printf("pid type is %T\n", pidEx)
	fmt.Printf("pid value is %v\n", pidEx)
	pidValue, _ := pidEx.(json.Number).Int64()
	fmt.Println("pidValue Int64 is ", pidValue)
}
