package main 

import "fmt"

func main (){
	
	var num int

	fmt.Scan(&num)

	if num >= 100 && num <= 999 {
		fmt.Print((num/10) % 10)
	
		} else {
		fmt.Print("Não possui 3 dígitos")
	}



}