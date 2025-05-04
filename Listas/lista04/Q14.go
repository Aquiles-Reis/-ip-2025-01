package main 

import "fmt"

func main (){
	
	var (
		vet1, vet2 [10]int
		vetres []int
		y int
	)
	fmt.Print("Digite os 10 valores do primeiro vetor:")
	for i:=0; i<10; i++{
		fmt.Scan(&vet1[i])
	}
	fmt.Print("Digite os 10 valores do segundo vetor:")
	for x:=0; x<10; x++{
		fmt.Scan(&vet2[x])
	}
	
	for _, z := range vet1 {
		vetres = append(vetres, z)
			vetres = append(vetres, vet2[y])
			y++
	}
	
	fmt.Print(vetres)
}