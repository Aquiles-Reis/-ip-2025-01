package main
import "fmt"
func main () {
	var salariomin , kw, consumo float32
	 fmt.Print("Digite o valor do salário mínimo:")
	 fmt.Scan(&salariomin)
	 fmt.Print("Digite a quantidade de kw gastos:")
	 fmt.Scan(&consumo)
	kw = salariomin * 0.007
	valorfinal:= consumo * kw
	desconto:= valorfinal * 0.9
	 fmt.Printf("O valor de cada kw é: %.2f\n", kw)
	 fmt.Printf("O valor do consumo será de: %.2f\n", valorfinal)
	 fmt.Printf("O valor com desconto será de %.2f\n:", desconto)
	
}