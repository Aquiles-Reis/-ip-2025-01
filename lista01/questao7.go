package main
import "fmt"
func main () {
	var F, pol float64
		fmt.Print("Digite o valor em Fahrenheit: ")
		fmt.Scan(&F)
		fmt.Print("Digite o valor da chuva em polegadas: ")
		fmt.Scan(&pol)
	C := (F - 32) * 5 / 9
	mm := pol * 25.4
		fmt.Printf("A temperatura em Celsius é: %.2f\n", C)
		fmt.Printf("O valor da chuva em milimetros é: %.2f\n", mm)
}