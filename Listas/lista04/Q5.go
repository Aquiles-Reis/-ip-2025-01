package main 

import "fmt"

func main (){

	var num[10]int

	fmt.Print("Digite 10 números:")
	
	for i:=0; i<10; i++{
	fmt.Scan(&num[i])
	}
	menor := num[0]
	posicao := 0
	for i := 1; i < 10; i++ {
		if num[i] < menor {
			menor = num[i]
			posicao = i
		}
	}
	fmt.Printf("O menor elemento do vetor é %d e sua posição dentro do vetor é: %d\n", menor, posicao )


	maior := num[0]
	local := 0
	for i := 1; i < 10; i++ {
		if num[i] > maior {
			maior = num[i]
			local = i
		}
	}
	fmt.Printf("O maior elemento do vetor é %d e sua posição dentro do vetor é: %d", maior, local  )
}