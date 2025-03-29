package main
import ("fmt"
        "math")
func main () {
	var h, a, v, Ab float64
	fmt.Print ("Digite o valor da altura da pirâmide hexagonal:")
	fmt.Scan (&h)
	fmt.Print ("Digite o valor do lado do hexágono:")
	fmt.Scan (&a)
	Ab = 3 * (math.Pow (a, 2) * math.Sqrt (3)) / 2
	v = (Ab * h) / 3
	fmt.Printf ("O volume da pirâmide hexagonal é: %.2f", v)
	
}