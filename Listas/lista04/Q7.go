package main 

import "fmt"

func main (){

	var num[]int

	for i:=1; i < 200; i += 2 {
		num = append(num, i)
	}
	fmt.Print(num)
}