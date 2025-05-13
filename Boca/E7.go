package main 

import "fmt"

func main (){

	var (
		n, soma int
	)
	
	fmt.Scan(&n)
	if n <= 0 {
		return
	}
	vetA := make ([]int, n)
	
	for z := 0; z < n; z++ {
		fmt.Scan(&vetA[z])
	}
	
	vetB := make ([]int, n)
	
	for y := 0; y < n; y++ {
		fmt.Scan(&vetB[y])
	}

	for i := 0; i < n; i ++ {
		soma += vetA[i] * vetB[i]
	}
	fmt.Print(soma)
}