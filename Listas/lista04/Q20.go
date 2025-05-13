package main 

import "fmt"

func main (){

	var (
		num [20] int
	)
	frequencia := make(map[int]int)
	
	for i := 0; i < 20; i ++ {
		fmt.Scan(&num[i])
	}
	for _, i := range num{
		frequencia [i] ++
	}
	for i := 1; i <= 6; i++ {
        fmt.Printf("Número %d: %d vezes\n", i, frequencia[i])
	}	
}