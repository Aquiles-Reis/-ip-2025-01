package main

import "fmt"

func soma (a , b float64) float64 {
		return a + b
}

 func main (){

	var num1, num2 float64

	fmt.Scan(&num1, &num2)

	resultado:= soma(num1, num2)

	fmt.Printf("A soma dos números é %f", resultado)
}
