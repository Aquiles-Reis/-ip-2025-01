package main 

import "fmt"

func main (){

var num[10]int
var soma, g int

for x := 0; x < 10; x++ {
	fmt.Scan(&num[x])
}
	for i, valorp := range num {
		if valorp %2 == 0 {
			i++
			fmt.Printf("%d ", valorp)
			soma = soma + valorp
		}
	}
	fmt.Println()
	fmt.Println(soma)

	for i, valori := range num {
		if valori %2 != 0 {
			i++
			fmt.Printf("%d ", valori)
			g++
		}
	}
	fmt.Println()
	fmt.Print(g)
		
}