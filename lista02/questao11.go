package main

import "fmt"

func main (){

	var x float64
	var f float64
	
	fmt.Print("Digite o valor de x:")
	fmt.Scan(&x)

	f = 8/(x-2)

	fmt.Print("O valor de f(x) é: ", f)

}