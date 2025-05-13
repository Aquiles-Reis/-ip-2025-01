package main 

import "fmt"

func main (){

	vetN := make([]int, 10)
	vetDiv := make([]int, 5)

	fmt.Println("Digite os valores do 1° vetor:")
	for i := range vetN {
		fmt.Scan(&vetN[i])
	}

	fmt.Println("Digite os valores do 2° vetor:")
	for i := range vetDiv {
		fmt.Scan(&vetDiv[i])
	}
	for _, x := range vetN {
		fmt.Printf("Número %d:\n", x)
		for m, n := range vetDiv {
			if x %n == 0 {
				fmt.Printf("       Divisível por %d na posição %d.\n", n, m)
			} 
		}
	}
	
}