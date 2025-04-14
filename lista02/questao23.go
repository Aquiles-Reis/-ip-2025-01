package main

import "fmt"

func main (){

	idades := []int{15, 17, 65}
	classes := []string{"Não-eleitor", "Eleitor facultativo", "Eleitor Obrigatório", "Eleitor facultativo"}
	
	var idade int

	fmt.Scan(&idade)

	classe := "Indefinido"
	for i, limite := range idades {
		if idade <= limite {
		   classe = classes[i]
		break 
		}		
	}
		if classe == "Indefinido" {
			classe = classes[len(classes)-1]
		}

		fmt.Print(classe)
}