package main

import (
	"fmt"
	"math"
)

func somatorio(n []float64) float64 {
	var soma float64
	for i:=0; i<50; i++ {
		soma += math.Pow((n[i] - n[99-i]), 3)
			
	}
	return soma
}
func main (){
	

	n := make([]float64, 100)
	
	for i:=0; i<100; i++ {
	fmt.Scan(&n[i])

}
	resultado := somatorio(n)
	fmt.Println(resultado)
}