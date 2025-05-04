package main 

import "fmt"

func main (){

	var n float64

	fmt.Print("Digite um valor:")
	fmt.Scan(&n)

	if n > 0 {
		fmt.Print("O número é positivo.")
		
		} else if n < 0 {
			fmt.Print("O valor é negativo.")
			
			} else {
				fmt.Print ("O valor é neutro.")}
}