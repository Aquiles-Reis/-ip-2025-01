package main

import ("fmt"
		"sort")

func main (){

	var g int
	num := make([]int, 10)

	for i:=0; i<10; i++ {
		fmt.Scan(&num[i])
	}
	sort.Ints(num)

	for x := 0; x < 9; x ++ {
		if num [x] == num[x+1] {
			g++
			} else {
				if g > 0 {
					fmt.Printf("O número %d repetiu %d vezes.\n", num[x], g+1)
				}
				g = 0
		}
	}
	if g > 0 {
		fmt.Printf("O número %d repetiu %d vezes.\n", num[9], g+1)
	}
}