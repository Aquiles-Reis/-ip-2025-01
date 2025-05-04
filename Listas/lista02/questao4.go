package main

import ("fmt"
		"math")

func main (){

	var n float64
	
	fmt.Print("Digite um número qualquer:")
	fmt.Scan(&n)

	if n >= 0 {
		fmt.Print(math.Sqrt(n))
	
		} else {
			fmt.Print(math.Pow(n, 2))
		}
}