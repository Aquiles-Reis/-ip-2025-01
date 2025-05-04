package main

import "fmt"

func main (){

	var n, ingr int
	var pop, ger, arqui, cad float64
	
	fmt.Print("Digite a quantidade de jogos:")
	fmt.Scan(&n)

	for i:= 0; i < n; i++ {

	fmt.Println("Informe a quantidade de pessoas que foram ao jogo:")
	fmt.Scan(&ingr)

	fmt.Println("Informe a porcentagem de pessoas respectivamente nas opções: popular, geral, arquibancada e cadeiras.")
	fmt.Scan(&pop, &ger, &arqui, &cad)

	renda:= float64(ingr) * ((pop*1) + (ger*5) + (arqui*10) + (cad*20))
		fmt.Printf("A renda foi de: %2.f\n", renda)
	
	}
}