package main
import "fmt"
func main () {
	var numero int
	fmt.Print("Digite um número:")
	fmt.Scan(&numero)
	if numero % 15 == 0 {
		fmt.Printf ("%v é divisível por 3 e 5.", numero)
		} else {
		fmt.Print("Não é divisível por 3 e 5.")
	}
}