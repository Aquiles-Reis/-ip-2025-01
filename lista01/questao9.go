package main
import "fmt"
func main () {
	var a, b, c, delta float32
	fmt.Print ("Digite o valor de suas variáveis:")
	fmt.Scan (&a, &b, &c)
	delta = (b*b) - 4*a*c
	fmt.Printf("%.2f\n", delta)
}