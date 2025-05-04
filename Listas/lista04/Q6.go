package main 

import "fmt"

func main (){

	var num[]int

	for i:=100; i >= 1; i--{
		num = append(num, i)
	}
	fmt.Print(num)
}