package main 

import "fmt"
		
func main (){

	var (
		nus [30] int
		vetres [] float64
	)
	fmt.Print("Digite 30 valores: ")
	
	for i:=0; i<30; i++ {
	fmt.Scan(&nus[i])
	}

	for i, x := range nus {
		if i %2 == 0 {
			vetres = append(vetres, float64(x) * 2)
		} else {
			vetres = append(vetres, float64(x) * 3)
		}
	}
	fmt.Print(vetres)
}