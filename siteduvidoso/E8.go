package main 

import ("fmt"
		"sort")

func main () {
	var (
		n int
	)
	
	fmt.Scan(&n)

	vet := make ([]int, n)

	for i := 0; i < n; i++{
		fmt.Scan(&vet[i])
	}
	sort.Ints(vet)
	
	dif := 0
	igual := 0
	
	for x := 0; x < n; x ++ {
		if x == 0 {
			if vet[x] != vet[x+1] {
				dif ++
			} else {
				igual ++
			}
		} else {	
		if vet[x] != vet[x-1] {
		dif ++
		} else {
			igual ++
	}
	}
}
	fmt.Print(dif)
}