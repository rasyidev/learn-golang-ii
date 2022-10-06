package tambahan

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestMarshallUnmarshallComplex(t *testing.T) {
	jsonString := `[
									{
										"key": "name",
										"value": "Rohmat"
									},
									{
										"key": "age",
										"value": 20
									}
								]
								`

	/* Expected:
	{
		"name": "Rohmat",
		"age": "20"
	}

	*/

	golangObject := []map[string]interface{}{}
	json.Unmarshal([]byte(jsonString), &golangObject)
	// fmt.Printf("%v\n", golangObject)

	result := map[string]interface{}{}

	for i := range golangObject {
		key := golangObject[i]["key"].(string)
		value := golangObject[i]["value"]

		result[key] = value
	}
	fmt.Println(result)

	resultString, err := json.Marshal(result)
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("RESULT STRING:\n %v \n", string(resultString))

}

/*
=== RUN   TestMarshallUnmarshallComplex
map[age:20 name:Rohmat]
RESULT STRING:
 {"age":20,"name":"Rohmat"}
--- PASS: TestMarshallUnmarshallComplex (0.00s)
PASS
ok      learn-go-ii/tambahan    1.269s
*/
