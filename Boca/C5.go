package main 

import "fmt"

func NumeroTriangular (n int32) bool {
	if n < 0 {
		return false 
	} 
	
	var i int32
	
	for i = 1; i * (i+1) * (i+2) <= n; i++ {
		if i * (i+1) * (i+2) == n {
			return true
			}
		}
		return false
}
 

func main(){

	var num int32
	
	fmt.Scan(&num)

	if NumeroTriangular(num) {
		fmt.Printf("%d eh triangular", num)
	} else {
		fmt.Printf("%d nao eh triangular", num)
	}
}