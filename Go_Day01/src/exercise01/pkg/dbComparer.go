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
	addedCakes(old_db, new_db)
	removedCakes(old_db, new_db)
	compareTime(old_db, new_db)
	compareIngridients(old_db, new_db)
}

func addedCakes(old_db readDB.DataBase, new_db readDB.DataBase){
	oldDBMap := make(map[string]readDB.Cake)
	for _, cake := range old_db.Cakes{
		oldDBMap[cake.Name] = cake
	}
	for _, newCake := range new_db.Cakes{
		_, exists := oldDBMap[newCake.Name]
		if !exists{
			fmt.Printf("ADDED cake \"%s\"\n", newCake.Name)
		}
	}
}

func removedCakes(old_db readDB.DataBase, new_db readDB.DataBase){
	newDBMap := make(map[string]readDB.Cake)
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

func compareTime(old_db readDB.DataBase, new_db readDB.DataBase){
	oldDBMap := make(map[string]readDB.Cake)
	for _, cake := range old_db.Cakes{
		oldDBMap[cake.Name] = cake
	}
	for _, newCake := range new_db.Cakes{

		oldCake, exists := oldDBMap[newCake.Name]
		if !exists{
			continue
		}
		if oldCake.Time != newCake.Time{
			fmt.Printf("CHANGED cooking time for cake \"%s\" - \"%s\" instead of \"%s\"\n", oldCake.Name, newCake.Time, oldCake.Time)
		}
	}
}

func compareIngridients(old_db readDB.DataBase, new_db readDB.DataBase){
	oldDBMap := make(map[string]readDB.Cake)
	for _, cake := range old_db.Cakes{
		oldDBMap[cake.Name] = cake
	}
	for _, newCake := range new_db.Cakes{
		oldCake, exists := oldDBMap[newCake.Name]
		if !exists{
			continue
		}
		oldIngredientsMap := make(map[string]readDB.Ingridient)
		for _, ingredient := range oldCake.Ingredients{
			oldIngredientsMap[ingredient.Ingredient_name] = ingredient
		}
		for _, newIngredient := range newCake.Ingredients {
			if _, exists := oldIngredientsMap[newIngredient.Ingredient_name]; !exists {
				// Новый ингредиент, которого нет в старом торте
				fmt.Printf("ADDED ingredient \"%s\" for cake \"%s\"\n", newIngredient.Ingredient_name, newCake.Name)
			}
		}
		for _, oldIngredient := range oldCake.Ingredients {
			if _, exists := oldIngredientsMap[oldIngredient.Ingredient_name]; !exists {
				// Старый ингредиент, который был удалён
				fmt.Printf("REMOVED ingredient \"%s\" for cake \"%s\"\n", oldIngredient.Ingredient_name, newCake.Name)
			}
		}
	}
}