package main

import "fmt"

func main() {

	ida := []float64{500.00, 350.00, 350.00, 300.00}
	idavolta := []float64{900.00, 650.00, 650.00, 550.00}

	var regiao, modelo int

	fmt.Scan(&regiao, &modelo)

	if modelo == 1 {
		fmt.Print(ida[regiao-1])

	} else if modelo == 2 {
		fmt.Print(idavolta[regiao-1])
	}
}
