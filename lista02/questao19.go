package main

import ("fmt"
		"math")
func main (){

	var opcao int
	var raio, alt float64
	pi := 3.14

	fmt.Scan(&opcao)

	if opcao == 1 {
		fmt.Print("Digite o raio e altura")
		fmt.Scan(&raio, &alt)
			vol := (pi * math.Pow(raio, 2) * alt) / 3
			area := pi * raio * math.Sqrt(math.Pow(raio,2) + math.Pow(alt, 2))
			
			fmt.Print (vol, area)
	
		} else if opcao == 2 {
			fmt.Print("Digite o raio e altura")
			fmt.Scan(&raio, &alt)
				vol := pi * math.Pow(raio, 2) * alt
				area := 2 * pi * raio * alt
			
				fmt.Print (vol, area)
			
			} else if opcao == 3 {
				fmt.Print("Digite o raio e altura")
				fmt.Scan(&raio, &alt)
				vol := (4/3) * pi * math.Pow(raio, 3)
				area := 4 * pi * math.Pow(raio, 2)

				fmt.Print (vol, area)
	} 	
}