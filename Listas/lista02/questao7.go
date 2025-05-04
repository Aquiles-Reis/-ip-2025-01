package main

import (
	"fmt"
)

func main() {
	var A, B, C int

	// Solicita os três valores ao usuário
	fmt.Print("Digite o primeiro valor inteiro: ")
	fmt.Scan(&A)
	fmt.Print("Digite o segundo valor inteiro: ")
	fmt.Scan(&B)
	fmt.Print("Digite o terceiro valor inteiro: ")
	fmt.Scan(&C)

	var MENOR, INTER, MAIOR int

	// Determina os valores
	if A < B && A < C {
		MENOR = A
		if B < C {
			INTER = B
			MAIOR = C
		} else {
			INTER = C
			MAIOR = B
		}
	} else if B < A && B < C {
		MENOR = B
		if A < C {
			INTER = A
			MAIOR = C
		} else {
			INTER = C
			MAIOR = A
		}
	} else {
		MENOR = C
		if A < B {
			INTER = A
			MAIOR = B
		} else {
			INTER = B
			MAIOR = A
		}
	}

	// Imprime os valores em ordem crescente
	fmt.Printf("\nOs valores em ordem crescente são: %d, %d, %d\n", MENOR, INTER, MAIOR)
}