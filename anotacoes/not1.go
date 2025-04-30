package main
import"fmt"
func main (){

	//declarar vetor e/ou slice:
	var n [100]float64 // vetor com obrigatoriamente 100 valores dentro

	x := make([]float64, 100) // slice com 100 valores predefinidos

	var z []float64 // slice vazio

	var y [0] float64 // vetor nulo
	fmt.Print(y)
}