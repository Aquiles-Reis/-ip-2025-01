package main 

import ("fmt"
		"math")

func main (){
	var (
		npar int
		val, soma float64
	)
	fmt.Scan(&npar)
	if npar %2 != 0 {
		return
	}
	
	nus := make ([]float64, npar)
	
	for i:=0; i < npar; i++{
		fmt.Scan(&nus[i])
	}
	for x := 0; x < npar/2; x++ {
		val = math.Pow((nus[x] - nus[(npar - 1) - x]), 3)
		soma += val
	}
	fmt.Printf("%.2f", soma)
}