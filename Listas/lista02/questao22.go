package main

import "fmt"

func main(){

	salminimo :=788.00
	valorhora := 10.00

	var matricula int
	var horasextra float64
	
	fmt.Scan(&matricula,&horasextra)

	salbruto := (3 * salminimo) + (valorhora * horasextra)

	salliquido := 0.68 * salbruto

	fmt.Printf("%.2f %.2f", salbruto, salliquido)
} 