package main

import "fmt"

func main(){

	var dia, mes, ano int
	meses:= []string{"jan", "fev", "mar", "abr", "mai", "jun", "jul", "ago", "set", "out", "nov", "dez"}

	fmt.Print ("Digite um dia, mês(1-12) e ano:")
	fmt.Scan(&dia, &mes, &ano)

	if mes>=1 && mes <=12 {
		fmt.Print("%d de %s de %d", dia, meses[mes-1], ano)
	} else {
		fmt.Print("Mês invalido")
		return
	}
}