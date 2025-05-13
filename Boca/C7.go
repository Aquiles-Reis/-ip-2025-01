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
	
	var n float64

	fmt.Scan(&n)

	const termos = 20
	var soma float64

	for i := 0; i < termos; i ++ {
		
		// calcula cada termo da série
		termo:= math.Pow(n, float64(i)) / float64(fatorial(i))

		//alterna entere somar e subtrair
		if i %2 != 0 {
			termo = -termo
		}
		soma += termo
	}
	soma = soma + n - 1 
	fmt.Print(soma)
}