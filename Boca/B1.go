package main 

import ("fmt"
		"math")
func main (){

	var A, B, C float64

	fmt.Scan(&A, &B, &C)

	if A == 0 {
		fmt.Print ("Nao e equacao do segundo grau")
	return }
	
	delta := math.Pow (B, 2) - 4*A*C

	if delta >= 0 {
		fmt.Printf("%.5f %.5f", (-1*B + math.Sqrt(delta))/(2*A), (-1*B - math.Sqrt(delta))/(2*A) ) 
	} else if delta < 0 {
		fmt.Print ("Nao ha raizes reais")
	}

}