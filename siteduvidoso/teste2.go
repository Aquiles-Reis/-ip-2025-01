package main

import (
	"fmt"
)

func main() {
	// Declaração das variáveis
	var pInicial, d, deltaP, pMin float64
	var qInicial, deltaQ int

	// Leitura da entrada
	fmt.Scan(&pInicial, &qInicial, &d, &deltaP, &deltaQ, &pMin)

	// Variáveis auxiliares
	preco := pInicial
	quantidade := qInicial
	lucroMaximo := -1.0
	precoMaximo := pInicial
	ingressosMaximo := qInicial

	// Cabeçalho da tabela
	fmt.Println("Preco Ingressos_Vendidos Lucro")

	// Loop para realizar as iterações enquanto o preço não for menor que pMin
	for preco >= pMin {
		// Calcula o lucro
		lucro := (preco * float64(quantidade)) - d

		// Exibe os valores na tabela
		fmt.Printf("%.2f %d %.2f\n", preco, quantidade, lucro)

		// Atualiza o lucro máximo
		if lucro > lucroMaximo {
			lucroMaximo = lucro
			precoMaximo = preco
			ingressosMaximo = quantidade
		}

		// Atualiza os valores para a próxima iteração
		preco -= deltaP
		quantidade += deltaQ
	}

	// Exibe o lucro máximo e os valores correspondentes
	fmt.Printf("Lucro maximo: %.2f\n", lucroMaximo)
	fmt.Printf("Na faixa de preco: %.2f com %d ingressos.\n", precoMaximo, ingressosMaximo)
}