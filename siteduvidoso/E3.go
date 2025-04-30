package main 

import "fmt"

func main (){
	
	var (
		n int
	)

	fmt.Scan(&n)
	nus := make ([]int, n)
	
	
	for i:=0; i < n; i++ {
		fmt.Scan(&nus[i])
	}

	for i := range nus {	
		if n == 1 {
			fmt.Print("0")
		} else if i == 0 {
			fmt.Printf("%d ", nus[i+1])
		} else if i == n - 1 {
			fmt.Printf("%d ", nus [i-1])
		} else {fmt.Printf("%d ", nus[i+1] + nus[i-1])
	}
}
}