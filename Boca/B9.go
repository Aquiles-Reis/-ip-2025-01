package main 

import ("fmt" 
		"math")

func main (){

	categorias:= []string{"Abaixo do peso", "Peso normal", "Sobrepeso", "Obesidade"}
	valores:= []float64{18.5 , 25 , 30}

	var peso, altura float64	
	
	fmt.Scan(&peso, &altura)

	IMC := peso/ math.Pow(altura, 2)

	categoria:= "Indefinido"
	for i, limite:= range valores {
		if IMC <= limite {
			categoria = categorias[i]
			break
		}
	} 
	if categoria == "Indefinido" {
			categoria = categorias [len(categorias)-1]
	}
			fmt.Print(categoria)
}
