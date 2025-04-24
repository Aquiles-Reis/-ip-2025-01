package main 

import "fmt"

func main (){

	var P_inicial, D, DeltaP, P_min, Lucro float64
	var Q_inicial, DeltaQ int
	LucroMaximo := -1.0
	Preco := 0.0
	ingressos := 0


		fmt.Scan(&P_inicial, &Q_inicial, &D, &DeltaP, &DeltaQ, &P_min)

		fmt.Println("Preco   Ingressos Vendidos    Lucro")
	for P_inicial >= P_min {

		Lucro = (P_inicial * float64(Q_inicial)) - D
		
		fmt.Printf(" %.2f           %d           %.2f\n",P_inicial, Q_inicial, Lucro)
		
		
		if Lucro > LucroMaximo {
			LucroMaximo = Lucro
			Preco = P_inicial
			ingressos = Q_inicial
		}
		
			Q_inicial = Q_inicial + DeltaQ
			P_inicial = P_inicial - DeltaP
	}
		fmt.Printf("Lucro maximo: %.2f\n", LucroMaximo)
		fmt.Printf("Na faixa de preco: %.2f com %d ingressos.", Preco, ingressos)
}