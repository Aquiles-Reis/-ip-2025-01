package main 

import "fmt"

func main (){
	
	var compra float64

	fmt.Print ("Digite o valor da compra:")
	fmt.Scan(&compra)

	if compra < 10 {
		fmt.Printf ("O lucro será de: %.2f.", compra*0.7)
		} else if 10 <= compra && compra < 30 {
			fmt.Printf ("O lucro será de: %.2f.", compra*0.5)
			} else if 30 <= compra && compra < 50 {
				fmt.Printf ("O lucro será de: %.2f.", compra*0.4)
				} else if compra >= 50 { 
					fmt.Printf ("O lucro será de: %.2f.", compra*0.3)
	}
}