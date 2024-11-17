package dbComparer

import (
	"Day01/exercise00/pkg/readDB"
	"flag"
	"fmt"
	"os"
)


func DbComparer(){
	filename_old, filename_new := getFileNames()
	old_DB := &readDB.DataBase{}
	old_DB.ReadXML(filename_old)
	new_DB := &readDB.DataBase{}
	new_DB.ReadJSON(filename_new)
	comparedb(*old_DB, *new_DB)
}

func getFileNames() (filename_old string, filename_new string){
	flag.StringVar(&filename_old, "old", "", "old_db")
	flag.StringVar(&filename_new, "new", "", "new_db")
	flag.Parse()
	if filename_old == "" || filename_new == ""{
		fmt.Printf("используй флаги --old или --new для указания путей файлов")
		os.Exit(1)
	}
	return
}

func comparedb(old_db readDB.DataBase, new_db readDB.DataBase){
	oldDBMap := make(map[string]readDB.Cake)
	newDBMap := make(map[string]readDB.Cake)
	for _, cake := range old_db.Cakes{
		oldDBMap[cake.Name] = cake
	}
	for _, newCake := range new_db.Cakes{
		_, exists := oldDBMap[newCake.Name]
		if !exists{
			fmt.Printf("ADDED cake \"%s\"\n", newCake.Name)
		}
	}
	for _, cake := range new_db.Cakes{
		newDBMap[cake.Name] = cake
	}
	for _, oldCake := range old_db.Cakes{
		_, exists := newDBMap[oldCake.Name]
		if !exists{
			fmt.Printf("REMOVED cake \"%s\"\n", oldCake.Name)
		}
	}
	
}