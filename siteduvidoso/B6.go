package main 

import "fmt"

func main (){

	var numconta int
	var lim, saldoin, dep, ret, saldofinal float64
	
	fmt.Scan(&numconta, &lim, &saldoin, &dep, &ret)

	saldofinal = saldoin + dep - ret

	if saldofinal <= lim {
		fmt.Printf("Saldo: %.5f", saldofinal)
		} else {
		fmt.Printf("Credito excedido: %.5f", saldofinal)
	}

}