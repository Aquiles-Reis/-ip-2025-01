package main 

import "fmt"

func main (){

	var (
		termo, x, n, pi float64
		g int
	)

	fmt.Scan(&n)
	if n < 1 || n > 1000 {
		return
	}
	for i := 1.0; i < 2*n; i = i + 2 {
		termo = (1/i)
		if g %2 != 0 {
			termo = -termo
		}
		x += termo
		g++
	}
	pi = 4 * x
	fmt.Printf("%.6f", pi)
}