package main

import "fmt"

func main (){

	var vet1[10]int
	var vet2[5]int
	var soma int
	var vetres1 []int
	var vetres2 []int

	
	fmt.Print("Digite os 10 valores do primeiro vetor:")
	
	for i:=0; i<10; i++{
		fmt.Scan(&vet1[i])
	}
	
	fmt.Print("Digite os 5 valores do segundo vetor:")
	
	for x:=0; x<5; x++ {
		fmt.Scan(&vet2[x])
	}

	for a, valor := range vet2 {
		soma = soma + valor
		a++
	}
	for b, par := range vet1 {
		if par %2 == 0 {
			vetres1 = append(vetres1, par + soma)
			b++
		}
	}
	fmt.Printf("Primeiro vetor resultante: %d\n", vetres1)

	for c, impar := range vet1 {
		if impar %2 != 0 {
			vetres2 = append(vetres2, impar + soma)
			c++
		}
	}
	fmt.Printf("Segundo vetor resultante: %d", vetres2)
}
