package main

import "fmt"

func main (){

	var conta int
	var consumidor string
	var consumo, valor float64
	
	fmt.Print("Informe, respectivamente, o número da sua conta, a quantidade de metros cúbicos usados e seu tipo de consumidor(R-residencial, C-comercial, I-industrial):")
	fmt.Scan(&conta, &consumo, &consumidor)

	if consumidor == ("R") { 
		valor = 5 + 0.05 * consumo
		
		fmt.Printf("Conta = %v\n", conta) 
		fmt.Printf("Valor da conta = %.2f", valor)
		
		} else if consumidor == ("C") {
			valor = 500 + 0.25 * (consumo - 80)
			
			fmt.Printf("Conta = %v\n", conta,)
			fmt.Printf("Valor da conta = %.2f\n", valor)
				
			} else if consumidor == ("I") {
					valor = 800 + 0.04 * (consumo - 100)
					
					fmt.Printf("Conta = %v\n", conta)
					fmt.Printf("Valor da conta = %.2f", valor)
					
					} else {
						fmt.Print("Tipo de consumidor inválido.")}
}