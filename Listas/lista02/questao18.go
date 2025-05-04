package main

import "fmt"

func main (){

	var dia int
	var categoria string
	var valor float64
	
	fmt.Scan(&valor)	
	fmt.Scan(&dia, &categoria)

	if dia == 2 || dia == 3 || dia == 5 && categoria == "normal" {
		fmt.Print(valor * 0.60)
	} else if categoria == "lançamento" {
		fmt.Print(valor *1.15)
	} else 
		{fmt.Print(valor)
	}
}