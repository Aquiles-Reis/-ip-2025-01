package main
import "fmt"
func main (){
	var n float64
	fmt.Print ("Digite um número:")
	fmt.Scan (&n)
	if n > 0 {
		fmt.Println ("O número é positivo.")
	}else if n < 0 {
		fmt.Println ("O número é negativo.")
	}else {
		fmt.Println ("O número é zero.")
	}
}