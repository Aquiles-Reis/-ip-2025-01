package main

import "fmt"

func main (){

	var n int
	
	fmt.Print ("Digite um número inteiro:")
	fmt.Scan (&n)

	if n % 5 == 0 {
		fmt.Print("O número é divisível por 5.")
		} else {
			fmt.Print("Onúmero não é divisível por 5.")
	} 
}