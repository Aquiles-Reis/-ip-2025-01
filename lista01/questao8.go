package main
import ("fmt"
		"math")
//m^2 = R$100
	func main (){
	var raio, altura float64
	pi := 3.14159
	 
	 fmt.Print("Digite o raio da base e a altura da lata:")
	 fmt.Scan(&raio, &altura)
	
	 area := (pi*math.Pow(raio, 2) *2) + (2*pi*raio*altura)
	 custo:= area * 100
	fmt.Printf("O custo da lata será de %.2f\n", custo)
}