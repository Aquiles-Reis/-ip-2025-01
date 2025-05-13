package main 

import ("fmt"
		"sort")


func main (){

	var
	(		
		num [] int
		nus int
	)

	for i := 0; i < 10; i++ {
		fmt.Scan(&nus)
		
		num = append(num, nus)
		sort.Ints(num)
		} 
		
		fmt.Println(num)
		
}