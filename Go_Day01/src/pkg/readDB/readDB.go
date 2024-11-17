package readDB

import (
	"log"
	"os"
)


type Ingridients struct{
	ingridient_name string
	ingridient_count float64
	ingridient_unit string
}

type Recipe struct{
	name string
	time float64
}

func parseJsonToRecipe(){
	dataFromDB,err := os.ReadFile("cake.json")
	if err != nil{
		log.Fatal(err)
		return
	}

	


	
	
}