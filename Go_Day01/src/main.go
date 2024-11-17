package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)




func main(){
	file, err := os.Open("cake.json")
	if err != nil{
		log.Fatal(err)
	}
	defer file.Close()	
	rd := bufio.NewReader(file)
	// buf := make([]byte, 10)
	// n, err := rd.Read(buf)
	// if err != nil && err != io.EOF{
	// 	fmt.Print("hui znaet")
	// }
	// fmt.Printf("прочитано %d байт: %s\n", n, buf)
	s, err := rd.ReadString('\n')
	fmt.Printf("%s\n", s)
}