package main

import (
	"fmt"
)

func calcularPrecoFinal(precoBase float64, opcoes []string) float64 {
	precosOpcoes := map[string]float64{
		"ar_condicionado":   1750.00,
		"pintura_metalica":  800.00,
		"vidro_eletrico":    1200.00,
		"direcao_hidraulica": 2000.00,
	}

	precoFinal := precoBase
	for _, opcao := range opcoes {
		precoFinal += precosOpcoes[opcao]
	}

	return precoFinal
}

func main() {
	var precoBase float64
	fmt.Print("Digite o preço inicial do carro: R$ ")
	fmt.Scan(&precoBase)

	opcoes := []string{}

	var resposta string
	fmt.Print("Ar condicionado (R$ 1750,00)? (sim/não): ")
	fmt.Scan(&resposta)
	if resposta == "sim" {
		opcoes = append(opcoes, "ar_condicionado")
	}

	fmt.Print("Pintura metálica (R$ 800,00)? (sim/não): ")
	fmt.Scan(&resposta)
	if resposta == "sim" {
		opcoes = append(opcoes, "pintura_metalica")
	}

	fmt.Print("Vidro elétrico (R$ 1200,00)? (sim/não): ")
	fmt.Scan(&resposta)
	if resposta == "sim" {
		opcoes = append(opcoes, "vidro_eletrico")
	}

	fmt.Print("Direção hidráulica (R$ 2000,00)? (sim/não): ")
	fmt.Scan(&resposta)
	if resposta == "sim" {
		opcoes = append(opcoes, "direcao_hidraulica")
	}

	precoFinal := calcularPrecoFinal(precoBase, opcoes)

	fmt.Printf("\nO preço final do carro é: R$ %.2f\n", precoFinal)
}