package main

import "fmt"

func main (){

	var l, r, numeros, soma int
	var media float64
	fmt.Scan(&l, &r)

	if l < 1 || r < 1 || l > 10000 || r > 10000 {
	return
	}
	if l %2 != 0 && r %2 != 0 && l == r {
		fmt.Println("0")
		fmt.Println("0")
		return
	} 
	if l %2 == 0 {
	for i:=l; i <= r; i += 2 {
		numeros++
		soma = soma + i
	}
	} else if l %2 != 0 {
	for i:=l+1; i <= r; i += 2 {
		numeros++
		soma = soma + i
	}
	}
	media = float64(soma) / float64(numeros)
	fmt.Println(soma)
	fmt.Print(media)
}