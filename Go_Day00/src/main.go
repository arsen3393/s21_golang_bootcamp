package main

import (
	"GoDay0/pkg/flagParser"
	"GoDay0/pkg/input"
	"GoDay0/pkg/output"
	"fmt"
)

func main(){
	UserFlags := flagParser.Parse()
	// fmt.Printf("%+v\n", UserFlags)
	numbers, err := input.Input()
	if err != nil{
		fmt.Print("Ошибка ввода")
	}else{
		output.Output(numbers, UserFlags)
	}
}
