package main
import "fmt"
func main (){
	var h, valor int
	fmt.Print ("Digite o número de horas estacionadas:")
	fmt.Scan(&h)
	
	valor= (h * 5) - (h/3*5)
		
	fmt.Print("Valor a ser pago: R$", valor)
}