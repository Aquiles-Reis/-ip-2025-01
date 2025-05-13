package main

import "fmt"

func main (){

	var (
		cod int
		num [10] int
	)

	fmt.Scan(&cod)
	if cod < 0 || cod > 2 {
		return
	}

	for i := 0; i < 10; i ++ {
		fmt.Scan(&num[i])
	}
	switch cod {
	case 0:
		break
	case 1:
		for _, x := range num {
			fmt.Printf("%d ", x)
		}
	case 2: 
		for z := 0; z < 10; z ++ {
			fmt.Printf("%d ", num[9 - z])
		}
	}

}