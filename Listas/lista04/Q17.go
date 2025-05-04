package main 

import "fmt"

	func primo(n int) bool{
		if n < 2 {
			return false
		}
		for i:=2; i*i <= n; i++ {
			if n %i == 0 {
			return false
			}
		}
		return true
	}
	
	func main (){

	var (
		nus[10]int
	)

	for i:=0; i<10; i++ {
		fmt.Scan(&nus[i])
	}
	for i, x := range nus {
		if primo(x) {
			fmt.Printf("Número: %d, Posição: %d \n", x, i)
		}
	}
}