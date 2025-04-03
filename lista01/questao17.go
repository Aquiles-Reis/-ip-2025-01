package main

import"fmt"

func main (){
	var x, y int
	fmt.Print("Digite o valor de x e y:")
	fmt.Scan(&x, &y)
	if x %2 == 1 {
		fmt.Print("Erro, o x é impar")
	
	} else {
		for i := 0; i < y; i++ {
		fmt.Print (x, " " )
		x = x + 2
	}
}

}