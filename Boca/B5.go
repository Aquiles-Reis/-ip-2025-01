package main

import "fmt"

func main (){

	var faltas int
	var n1, n2, n3, media float64
	
	fmt.Scan(&n1, &n2, &n3, &faltas)

	if faltas > 16 {
		fmt.Print ("Reprovado por falta")
		return}

	media = (n1 + n2 + n3)/3

	if media >= 7 {
		fmt.Print("Aprovado")
	} else if media <= 7 && media >= 5 {
		fmt.Print ("Prova final")
	} else {
		fmt.Print("Reprovado")
	}

}