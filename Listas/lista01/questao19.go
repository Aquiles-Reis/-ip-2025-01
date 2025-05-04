package main

import "fmt"

func main (){
	var n, i int
	var S, r float64
	S = 1
	fmt.Print("Digite o valor de n:")
	fmt.Scan(&n)

	if n < 2 {
		fmt.Print("O valor de n tem de ser inteiro e maior que 1.")
	} else { 
		for i=2; i < n+1; i++ {

			r = 1/float64(i)
			S = S + r
		}
		fmt.Printf("%.6f", S)

	}

}