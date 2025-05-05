package main 

import "fmt"

func main (){

	var (
		n int
	)
	fmt.Scan(&n)
	if n <= 0 {
		return
	}
	vet := make ([]int, n)
	for i := 0; i < n; i ++ {
		fmt.Scan(&vet[i])
	}
	if n %2 == 0 {
		for i := 0; i < n/2; i ++ {
			if vet[i] != vet [n-1-i] {
				fmt.Println("Não")
				return
			}
		}
		fmt.Println("Sim")	
	}
}