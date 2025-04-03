package main

import "fmt"

func main (){
	var matriz [2][2] float64
	var determinante float64

	fmt.Print("Digite os valores da matriz 2x2 (a, b, c, d): ")
	fmt.Scan(&matriz[0][0], &matriz[0][1], &matriz[1][0], &matriz[1][1])
	determinante = matriz[0][0]*matriz[1][1] - matriz[0][1]*matriz[1][0]
	fmt.Printf("O determinante da matriz é: %.2f\n", determinante)

}