package models

import	(
	"fmt"
	"os"
//	"io/ioutil"
//	"encoding/json"
)

func ErrCheck(err error) {
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}

type models struct {
    Id int
    Name string
    Value string
    Date string
}


func DataBaseList() {}
func DataBaseAdd() {}
/*
func DataBaseMarshal(text []byte) textJson{
	structJson := models {
		Id: 0,
		Name: "Some name",
		Value: "Some value",
		Date: "Some time",
	}
	textJson, _ := json.MarshalIndent(structJson, text)
	return textJson
}
*/
func DataBaseUnmarshal() {}

func DataBaseLoad(text string) {
	file, err := os.OpenFile("intermal/bd/models_test.bd", os.O_APPEND|os.O_CREATE, 0600)
	defer file.Close()
	ErrCheck(err)
	file.WriteString(text)
}
