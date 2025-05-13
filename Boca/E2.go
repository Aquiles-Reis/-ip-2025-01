package main 

import ("fmt"
		"sort")

func main (){

	var (
		n int
	x float64
	)

	fmt.Scan(&n)
	if n <= 0 {
		return
	}	
	nus := make ([]float64, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&nus[i])
	}
	sort.Float64s(nus)

	
	if nus[n-1] == nus[0] {
		for z := 0; z < n; z++ {
			fmt.Print("0.00 ")
		}
		} else {
			for y:=0; y < n; y++{
			x = (nus[y] - nus[0]) / (nus[n-1] - nus[0])
			fmt.Printf("%.2f ", x)
			}
		}
}