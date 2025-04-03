package main

import "fmt"

func main (){
	var h, min, s int
	fmt.Print ("Digite as horas, os minutos e os segundos:")
	fmt.Scan(&h, &min, &s)
	tempo:= h*3600 + min*60 + s
	fmt.Printf ("O tempo em segundos é: %v\n", tempo)
}