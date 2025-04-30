package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Print("Digite o número de elementos: ")
	fmt.Scan(&N)

	numeros := make([]int, N)
	fmt.Println("Digite os números:")
	for i := 0; i < N; i++ {
		fmt.Scan(&numeros[i])
	}

	// Variáveis para armazenar o tamanho do maior trecho crescente
	maxLen := 1
	currLen := 1

	for i := 1; i < N; i++ {
		if numeros[i] > numeros[i-1] {
			currLen++
			if currLen > maxLen {
				maxLen = currLen
			}
		} else {
			currLen = 1
		}
	}

	fmt.Printf("O tamanho do maior trecho consecutivo estritamente crescente é: %d\n", maxLen)
}