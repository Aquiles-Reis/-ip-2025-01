package main 

import "fmt"
		

func main (){

	var B, E int
	
	fmt.Scan (&B, &E)

	x := B

	if E == 0 {
		fmt.Print(1)
		} else if E != 0 {
			for i := 1; i < E; i++ {
			B = B * x
		}
	fmt.Print(B)
}
}