package main 

import "fmt"

func main (){

	var codigo int
	var valor, preco float64
	
	fmt.Scan(&valor)
	fmt.Scan(&codigo)


	if codigo == 1 {
		preco = valor * 0.9
		
		} else if codigo == 2 {
			preco = valor * 0.95
		
			} else if codigo == 3 {
				preco = valor
				
				} else if codigo == 4 {
					preco = valor * 1.1
	}
		
		fmt.Printf("%f", preco)
}