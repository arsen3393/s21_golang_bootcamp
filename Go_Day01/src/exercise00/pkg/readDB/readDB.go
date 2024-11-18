package readDB

import (
	"encoding/xml"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
)


type Ingridient struct{
	Ingredient_name string `json:"ingredient_name" xml:"itemname"`
	Ingredient_count string `json:"ingredient_count" xml:"itemcount"`
	Ingredient_unit string `json:"ingredient_unit" xml:"itemunit"`
}

type Cake struct{
	Name string `json:"name" xml:"name"`
	Time string `json:"time" xml:"stovetime"`
	Ingredients []Ingridient `json:"ingredients" xml:"ingredients"`
}

type DataBase struct{
	XMLName xml.Name `xml:"recipes"`
	Cakes [] Cake `json:"cake" xml:"cake"`
}

func ReadDB(){
	filename, format := getFileName()
	if format == "json"{
		outputJSON(filename)
	}else if format == "xml"{
		outputXML(filename)
	}else{
		fmt.Print("Неправильный формат файла, используй xml или json\n")
		os.Exit(1)
	}
}

func getFileName() (filename string, format string){
	format = ""
	flag.StringVar(&filename, "f","","flag string")
	flag.Parse()
	if filename == ""{
		fmt.Printf("Error: используй флаг -f для пути к файлу")
		os.Exit(1)
	}
	extenction := filepath.Ext(filename)
	if extenction == ".json"{
		format = "json"
	}else if  extenction == ".xml" {
		format = "xml"
	}
	return
}

func getDataFromFile(filename string) (dataFromFile []byte){
	dataFromFile, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	return
}
