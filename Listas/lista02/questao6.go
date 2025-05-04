package main 

import "fmt"

func main (){

	var a, b int

	fmt.Print("Digite dois valores:")
	fmt.Scan(&a, &b)

	if a % b == 0 {
		fmt.Print(a, " é divisível por ", b, ".")
		} else {
			fmt.Print(a, " não é divisível por ", b, ".")
		}
}