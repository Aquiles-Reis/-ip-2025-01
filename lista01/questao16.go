package main

import "fmt"

func main () {
	var salario float64
	fmt.Print("Digite o salário:")
	fmt.Scan(&salario)
	if salario <= 300 {
		reajuste := salario * 1.5
		fmt.Printf("O salário com reajuste será de %.2f\n", reajuste)
		} else {
			reajuste := salario * 1.3
			fmt.Printf("Osalário com reajuste será de %.2f\n", reajuste)
		}
	
}