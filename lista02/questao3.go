package main

import "fmt"

func main (){

	var a, b, s int

		fmt.Print("Digite dois valores inteiros:")
		fmt.Scan(&a, &b)

	s = a + b

		if s > 20 {
			fmt.Print(s + 8)
		
			} else if s <= 20 {
				fmt.Print(s - 5)}
}