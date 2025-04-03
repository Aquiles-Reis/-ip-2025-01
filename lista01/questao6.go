package main

import "fmt"

func main (){
	var F float64
	var n, i int
	fmt.Print("Digite o número de temperaturas a serem convertidas: ")
	fmt.Scan(&n)
	
	
	for i = 0 ; i < n; i++ {
		fmt.Print("Digite a temperatura em Fahrenheit: ")
		fmt.Scan(&F)
	
		C := (F - 32) * 5 / 9
		fmt.Printf("A temperatura em Celsius é: %.2f\n", C)
	}
}