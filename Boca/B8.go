package main 

import "fmt"

func main (){
	
	var venda,salariobase, salariofinal float64

	salariobase = 500

	fmt.Scan(&venda)

	if venda <= 15000 {
		salariofinal = salariobase + (0.09 * venda)
		fmt.Printf("%.5f", salariofinal)

	} else {
		salariofinal = salariobase + 800 + (0.09 * venda)
		fmt.Printf("%.5f", salariofinal)
	}

}