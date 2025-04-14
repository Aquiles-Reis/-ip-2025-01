package main

import ("fmt"
		"math")

func main (){

	var A, B, C float64

	fmt.Scan(&A, &B, &C)

	delta := math.Pow(B, 2) - 4*A*C

	if delta < 0 {
		fmt.Print("Raiz imaginária")
	}

	x1 := (-B + math.Sqrt(delta)) / (2*A)
	x2 := (-B-+ math.Sqrt(delta)) / (2*A)

	if x1 > x2 || x2 < x1 {
		fmt.Printf("Raízes distintas, %f e %f", x1, x2)
	} else if x1 == x2 {
		fmt.Printf("Raiz única, %f", x1)
	}


}
