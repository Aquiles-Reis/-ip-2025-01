package main

import "fmt"

func main (){

	var notas[15]int

	for i := 0; i < 15; i ++ {
		fmt.Scan(&notas[i])
	} 
	frequenciaAbsoluta := make(map[int]int)
	totalNotas := len (notas)

	for _, nota:= range notas {
		frequenciaAbsoluta [nota]++
	}

	fmt.Println("Nota | Frequência Absoluta | Frequência Relativa")
	fmt.Println("-------------------------------------------------")
	for nota:=0; nota <= 10; nota++ {
		fa := frequenciaAbsoluta[nota]
		fr := float64(fa) / float64(totalNotas)
		fmt.Printf("%4d | %19d | %18.2f \n", nota, fa, fr)
	}

}