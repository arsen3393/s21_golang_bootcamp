package readDB

import (
	"encoding/json"
	"fmt"
	"log"
)

func (db *DataBase) ReadJSON(filename string){
	dataFromFile := getDataFromFile(filename)
	if err := json.Unmarshal(dataFromFile, &db); err != nil{
		log.Fatal(err)
	}

}


func outputJSON(filename string){
	db := &DataBase{}
	db.ReadJSON(filename)
	data, err := json.MarshalIndent(&db, "", "    ")
	if err != nil{
		log.Fatal(err)
	}
	fmt.Printf("%s", data)
}
