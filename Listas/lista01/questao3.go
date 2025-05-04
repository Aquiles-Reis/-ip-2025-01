package main

import (
	"fmt"
	"strconv")
func main () {
	
	var n1, n2, n3 string 
	fmt.Print("Digite o valor de n1, n2 e n3:")
	fmt.Scan(&n1, &n2, &n3)
	
	a := n1 + n2 + n3
	
	b, err:= strconv.Atoi (a)
		if err != nil {
			fmt.Println("Erro: ", err)
		} else {
			fmt.Println(a,",",b*b)}
	

}