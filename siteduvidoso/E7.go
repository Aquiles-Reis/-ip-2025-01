package main 

import "fmt"

func main (){

	var (
		n, soma int
	)
	fmt.Print("Digite a quantidade de números nos vetores:")
	fmt.Scan(&n)
	if n <= 0 {
		return
	}
	vetA := make ([]int, n)
	fmt.Print("Digite os números inteiros do 1° vetor:")
	for z := 0; z < n; z++ {
		fmt.Scan(&vetA[z])
	}
	
	vetB := make ([]int, n)
	fmt.Print("Digite os números inteiros do 2° vetor:")
	for y := 0; y < n; y++ {
		fmt.Scan(&vetB[y])
	}

	for i := 0; i < n; i ++ {
		soma += vetA[i] * vetB[i]
	}
	fmt.Print(soma)
}