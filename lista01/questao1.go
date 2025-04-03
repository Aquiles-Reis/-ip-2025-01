package main

import "fmt"

func main () {
	var n1, n2, n3 float32
	var media float32
	fmt.Print ("Adicione as 3 notas:")
	fmt.Scan (&n1, &n2, &n3)
	media = (n1 + n2 + n3) / 3
	if media >= 6 {
		fmt.Print ("Aprovado" ,  media)
		} else {
			fmt.Print ("Reprovado " ,  media)
		}
}