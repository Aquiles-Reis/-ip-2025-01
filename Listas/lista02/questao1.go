package main

import "fmt"

func main(){
	
	var n int

	fmt.Print("Digite um valor inteiro positivo:")
	fmt.Scan(&n)

	if n %2==0 {
		fmt.Print ("O número é par.")
	
		} else {
			fmt.Print ("O número é ímpar.")
		}
}