package main 

import "fmt"

func main (){

	var nota int

	fmt.Scan(&nota)
	if nota < 0 || nota > 10 {
		return
	}

	if nota < 3 {
		fmt.Print("E")
	} else if nota < 5 {
		fmt.Print("D")
	} else if nota < 7 {
		fmt.Print("C")
	} else if nota < 9 {
		fmt.Print("B")
	} else {
		fmt.Print("A")
	}

}