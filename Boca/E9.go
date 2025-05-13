package main 

import "fmt"

func main (){

	var (
		n, soma int
		vetF [] int
	)

	fmt.Scan(&n)
	if n <= 0 {
		return
	}

	vetM := make ([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&vetM[i])
	}
	
	for i,_ := range vetM {
		soma += vetM[i]
		vetF = append(vetF, soma)
	}
	for _, x := range vetF {
		fmt.Printf("%d ", x)
	}
}