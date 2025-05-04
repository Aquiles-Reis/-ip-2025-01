package main 

import "fmt"

func crescente()
	

func main (){

	var (
		n int
	)

	fmt.Scan(&n)
	if n <= 0 {
		return
	}
	
	nus := make([]int, n)

	for i:=0; i < n; i++{
		fmt.Scan(&nus[i])
	}


}