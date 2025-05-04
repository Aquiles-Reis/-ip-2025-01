package main 

import "fmt"

func main (){

	var ns [10] int

	for i:=0; i < 10; i ++ {
		fmt.Scan(&ns[i])
	}
	fmt.Println("Os valores mariores que cinquenta encontrados e suas respectivas posições são:")
	
	encontrou := false
	for i, num := range ns {
		if num > 50 {
			fmt.Printf("Número: %d | Posição: %d\n", num, i)
			encontrou = true
		}
	}
	if ! encontrou {
		fmt.Println("Nenhum valor maior que 50 foi encontrado")
	}

}