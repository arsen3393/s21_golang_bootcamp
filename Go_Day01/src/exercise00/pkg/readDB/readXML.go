package readDB

import (
	"encoding/xml"
	"fmt"
	"log"
)


func (db *DataBase) ReadXML(filename string){
	dataFromFile := getDataFromFile(filename)
	if err := xml.Unmarshal(dataFromFile, &db); err != nil{
		log.Fatal(err)
	}
}

func outputXML(filename string){
	db := &DataBase{}
	db.ReadXML(filename)
	data, err := xml.MarshalIndent(&db, "", "    ")
	if err != nil{
		log.Fatal(err)
	}
	fmt.Printf("%s", data)
}