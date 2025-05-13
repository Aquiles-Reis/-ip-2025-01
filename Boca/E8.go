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
	
	dif := 1
	
	for x := 1; x < n; x ++ {
		if vet[x] != vet[x-1] {
			dif ++
				}
			}	
	fmt.Print(dif)
}