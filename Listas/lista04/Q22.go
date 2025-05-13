package main

import ("fmt"
		"time")

func main (){

	var (
		total [20] float64
		conta [] int
		saldo [] float64
		cod, numconta int
		val float64
	)

	fmt.Println("Digite o código da conta seguido do seu saldo.")
	for i := 1; i <= 4; i ++ {
		fmt.Scan(&total[i])
	}
	for x, y := range total {
		if x %2 != 0 {
			conta = append(conta, int(y))
		} else {
			saldo = append(saldo, y)
		}
	}
	// menu principal.
	for {
	fmt.Print("\n1. Efetuar depósito.\n")
	fmt.Print("2. Efetuar saque.\n")
	fmt.Print("3. Consultar o ativo bancário.\n")
	fmt.Print("4. Finalizar o programa.\n")
	fmt.Print("\nDigite o número da opção acima desejada.\n")
	fmt.Scan(&cod)

	// fim do menu.
	if cod == 4 {
		fmt.Print("Programa finalizado.")
		break
	}
		
	encontrou := false
	switch cod {
	
	// Efetuar um depósito.
	case 1:
		fmt.Print("Digite o código da conta.\n")
		fmt.Scan(&numconta)
		for _, n := range conta {
			if numconta == n {

				
				fmt.Print("Digite o valor do depósito.\n")
				encontrou = true
				
				fmt.Scan(&val)
				
				saldo[n] += val	
				fmt.Printf("\nO seu saldo final é de: R$%.2f.\n", saldo[n])
				time.Sleep(1 * time.Second)
			}
		}
		if !encontrou {
			fmt.Print("Conta não encontrada.")
			time.Sleep(1 * time.Second)
		
		}
		
	// Efetuar um saque.
	case 2:
		fmt.Print("Digite o código da conta.\n")
		fmt.Scan(&numconta)
		for _, n := range conta {
			if numconta == n {

				fmt.Print("Digite o valor do saque.\n")
				encontrou = true
				
				fmt.Scan(&val)
				
				// Verificando se o saldo é suficiente.
				if val > saldo[n] {
					fmt.Print("\nSaldo insuficiente.\n")
				} else {
				saldo[n] -= val	
				fmt.Printf("\nO seu saldo final é de: R$%.2f.\n", saldo[n])
				time.Sleep(1 * time.Second)
				}
			}
		}

		if !encontrou {
			fmt.Print("Conta não encontrada.")
			time.Sleep(1 * time.Second)
			}
		
	// Consultar o ativo bancário
	case 3:
		var saldototal float64 
		
		for _, g := range saldo {
			saldototal += g
		}	
		fmt.Printf("O ativo bancário é de R$%.2f", saldototal)
		time.Sleep(1 * time.Second)
	}
	} 
}