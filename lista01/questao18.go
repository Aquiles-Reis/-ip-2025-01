package main

import "fmt"

func main (){

	var a, r, n, i, soma float64

	fmt.Print ("Digite o valor incial, a razão e a quantidade de termos da p.a.:")
	fmt.Scan (&a, &r, &n)
		
		for i = 0; i < n; i++ {
		soma = soma + a
		a = a + r	
		fmt.Println(a)
		}
		fmt.Print(soma)
}