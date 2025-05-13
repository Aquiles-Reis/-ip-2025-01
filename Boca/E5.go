package main 

import "fmt"

func main (){

	var (
		n int
	)

	fmt.Scan(&n)

	nus := make([]int, n)
	for i:=0; i < n; i++ {
		fmt.Scan(&nus[i])
	}
	maxLen := 1
	curLen := 1


	for i := 1; i < n; i ++ {
		if nus[i] > nus [i - 1] {
			curLen ++
			if curLen > maxLen {
			maxLen = curLen 
			}
		} else { 
			curLen = 1
		}
	}
	fmt.Print(maxLen)
}