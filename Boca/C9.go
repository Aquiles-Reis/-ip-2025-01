package main

import ("fmt"
		"math")

func fatorial (n int) int {
	if n == 0 || n == 1 {
		return 1
	} 
	fat := 1
	for i := 2; i <= n; i ++ {
		fat *= i
	}
	return fat
}

func main (){

	var x float64
	var soma float64
	const termos = 20

	fmt.Scan(&x)

	for i := 0; i < termos; i += 2 {
		
		termo := (math.Pow(x, float64(i))) / float64(fatorial(i))
		if i %4 == 2 {
			termo = -termo
		}
		soma += termo
	} 
	cos := math.Cos(x)
	diferenca := soma - cos
	fmt.Printf("%.4f %.4f %.4f", soma, cos, diferenca)
}