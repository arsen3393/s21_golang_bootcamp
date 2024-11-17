package readDB

import (
	"encoding/xml"
	"fmt"
	"log"
)


func readXML(filename string){
	db := new(DataBase)
	dataFromFile := getDataFromFile(filename)
	if err := xml.Unmarshal(dataFromFile, &db); err != nil{
		log.Fatal(err)
	}
	outputXML(*db)
}

func outputXML(db DataBase){
	data, err := xml.MarshalIndent(&db, "", "    ")
	if err != nil{
		log.Fatal(err)
	}
	fmt.Printf("%s", data)
}