package main

import ("fmt"
		"sort"
)

func main () {

	n := 3
	numeros:= make([]int ,n) 
	
	
	for i := 0; i < n; i++ {
	fmt.Scan(&numeros[i])
	}

		sort.Ints(numeros)

	for _, numero := range numeros {
 	fmt.Printf("%d ", numero )
	}
}