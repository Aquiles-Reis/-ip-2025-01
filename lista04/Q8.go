package main 

import ("fmt"
		"math")
func main (){

	var num[15] int
	var val[]float64

	for i:=0; i < 15; i++{
	fmt.Scan(&num[i])
	} 
	for _, raiz := range num{
	if raiz < 0 {
		val = append(val, -1)
	} else {
		g := math.Sqrt(float64(raiz))
		val = append(val, g)
		}
	}
	fmt.Printf("%.2f", val)

}