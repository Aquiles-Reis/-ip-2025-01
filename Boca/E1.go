package main 

import "fmt"

func main (){

	var (
		nAtle int
		soma, alt float64
	)

	fmt.Scan(&nAtle)
	alturas := make ([]float64, nAtle)

	for i:=0; i < nAtle; i++{
		fmt.Scan(&alturas[i])
	} 
	for _, alt = range alturas {
		soma += alt
	}
	media := soma / float64(nAtle)
	for _, alt = range alturas {
	if alt > media {
		fmt.Printf("%.2f\n", alt)
	}
}
}