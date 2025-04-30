package main 

import ("fmt"
		"sort")

func main (){

	var (
		idades [] int
		g int
		moda []int 
	)


	for i:=0; i < 4; i ++ {
		fmt.Scan(&idades[i])
	}
	sort.Ints(idades)	
	
	for i, ids := range idades {
		
		if idades[i] == idades [i - 1] {
			g ++
			moda = append(moda, ids)
		}
	}
	fmt.Printf("A moda é %d e aparece %d vezes.", moda, g)

}