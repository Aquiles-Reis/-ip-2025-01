package main

import "fmt"

func main (){

	var X, Y, Z int

	fmt.Scan(&X, &Y, &Z)

	if X + Y <= Z || X + Z <= Y || Y + Z <= X {
		fmt.Print("Nao forma triangulo")
	} else if X == Y && Y == Z {
		fmt.Print("Equilatero")
	} else if X == Y || Y == Z || Z == 0 {
		fmt.Print("Isosceles")
	} else {
		fmt.Print("Escaleno")
	}

}