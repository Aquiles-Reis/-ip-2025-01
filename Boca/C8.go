package main

import ("fmt"
		"math")

func main (){

	var x, sinA float64
	fmt.Scan(&x)

	for i := 0.0; i <= x + 0.05; i += 0.1 {
		sinA = i - math.Pow(i, 3)/6 + math.Pow(i, 5)/120 - math.Pow(i, 7)/(5040)
		fmt.Printf("%.1f %.4f\n", i, sinA)
	}
}