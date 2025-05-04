package main

import "fmt"

func main (){

	var nota1, nota2, x float64
	fmt.Scan(&x)
	
	alunos := [][]float64{}

	for i:= 0.0; i < x; i++{
		fmt.Print("Insira as duas notas: ")
		fmt.Scan(&nota1, &nota2)

		alunos = append(alunos, []float64{nota1, nota2})
	}

	fmt.Print(alunos)


	

}