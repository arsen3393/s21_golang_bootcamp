package readDB

import (
	"encoding/json"
	"fmt"
	"log"
)

func readJSON(filename string){
	db := new(DataBase)
	dataFromFile := getDataFromFile(filename)
	if err := json.Unmarshal(dataFromFile, &db); err != nil{
		log.Fatal(err)
	}
	outputJSON(*db)
	//output
}

func outputJSON(db DataBase){
	data, err := json.MarshalIndent(&db, "", "    ")
	if err != nil{
		log.Fatal(err)
	}
	fmt.Printf("%s", data)
}