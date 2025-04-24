package main

import "fmt"

func main (){

	var N int
	
	fmt.Scan(&N)

	if N %2 == 0 {
		fmt.Printf("%d Par", N)
	} else {
		fmt.Printf("%d Impar", N)
	}

}