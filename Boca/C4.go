package main

import "fmt"

func main (){

	var n int
	stopTerm := 0 
	
	for { 
	fmt.Scan(&n)
	if n == stopTerm {
	break	
	}
	validar := false
	for i:=1; i * i <= n; i++ {
		if n == i * i {
			fmt.Printf("%d eh quadrado perfeito\n", n)
			validar = true
			break
			} 
		}
		if ! validar {
			fmt.Printf("%d nao eh quadrado perfeito\n", n)
		}
	}
}
