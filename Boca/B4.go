package main 

import "fmt"

func main (){

	categorias:= []string{"Infantil A", "Infantil B", "Juvenil A", "Juvenil B", "Adulto"}
	idades:= []int{7, 10, 13, 17, 18}
	var idade int

	fmt.Scan(&idade)

	categoria:= "Indefinido"
	for i, limite := range idades {
		if idade <= limite {
			categoria = categorias [i]
			break
		}
		if categoria == "Indefinido" {
		categoria = categorias [len(categorias)-1]
		}
	}
		fmt.Print(categoria)


}