package main 

import "fmt"

func main (){

	var alt[10] float64
	var soma, med float64
	var maior[]float64
	
	for i:=0; i<10; i++ {
		fmt.Scan(&alt[i])
	}
	for _, med = range alt {
		maior = append(maior, med)
		soma += med
	}
	media := soma / 10

	for _, x := range maior {
		if x > media {
		
		fmt.Println(x)
	}

}
}