package models

import	(
	"os"
	"io/ioutil"
	"encoding/json"
)

func ErrCheck(err error) {
	if err != nil {
		fmt.Pritnln("Error:", err)
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

func DataBaseMarshal(text []byte) text  {
	structJson := models{
		id: 0
		name: "Some name",
		value: "Some value",
		timestamp: "Some time"
	}
	textJson, _ := json.MarshalIndent(structJson, text)
	return textJson
}

func DataBaseUnmarshal() {}

func DataBaseLoad(writer string) {
	file, err := os.Open("module/Service/intermal/bd/mosels.bd")
	defer file.Close()
	file, err := ioutil.ReadFile("module/Service/intermal/bd/models.bd")
	ErrCheck(err)
	file.Write(writer)
	file.Sync()
}
