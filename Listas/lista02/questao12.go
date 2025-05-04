package main

import "fmt"

func main (){

	var idade int
	
	fmt.Print ("Digite sua idade:")
	fmt.Scan (&idade)

	if idade >= 0 && idade <= 2 {
		fmt.Print("Recém-nascido")
		} else if idade >= 3 && idade <= 11 {
			fmt.Print("Criança")
			} else if idade >= 12 && idade <= 19 {
				fmt.Print("Adolescente")
				} else if idade >= 20 && idade <= 55 {
					fmt.Print("Adulto")
					} else if idade >= 56 {
						fmt.Print("Idoso")
						} else {
							fmt.Print("Idade inválida")
	}
}