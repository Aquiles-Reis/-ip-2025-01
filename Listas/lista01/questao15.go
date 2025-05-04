package main

import "fmt"

func main(){

	var i, x float64
	var N int

	fmt.Print("Digite um valor entre 5 e 2000:")
	fmt.Scan(&N)

	if N %2 == 0 { 
		for i=0; i < float64(N/2); i++ {
			x = x + 2
			fmt.Print (x*x, "\n")}
		} else {
			for i=0; i < float64(N-1)/2; i++ {
				x = x + 2
				fmt.Print (x*x, "\n")
				
		}
	}
}