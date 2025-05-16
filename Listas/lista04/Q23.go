package main

import ("fmt"
		"os"
   		"os/exec"
		"time"
		"sort"
	)
		
func clearTerminal() {
    cmd := exec.Command("cmd", "/c", "cls")
    cmd.Stdout = os.Stdout
    cmd.Run()
}

func completo(poltronas []int) {
	sort.Ints(poltronas)
		ocupadas := true

	for _, assento := poltronas {
		if assento == 0 {
			ocupadas = false
			break
		}
	}
	if ocupadas {
		fmt.Print("Todas as poltronas estão ocupadas.")
		time.Sleep(1 * time.Second)
	}

}

func main (){
	
	var (
		opcao int
	)
	// Criação das posições vazias no ônibus, janela e meio.
	jan := make([]int, 24)
	meio := make([]int, 24)

	for {
	fmt.Print("Você deseja sentar na janela ou no meio?\nDigite 1 para 'janela' ou 2 para 'meio'.\n")
	fmt.Scan(&opcao)

	switch opcao {
	case 1:
		var n int
		
		fmt.Print("\nDigite o número da poltrona na qual você deseja sentar.\n")
		fmt.Scan(&n)

		if n < 1 || n > 24 {
			fmt.Println("Esta poltrona não existe, escolha outra entre 1 e 24.")
			continue
		}

		if jan[n-1] == 1 {
			fmt.Print("Esta poltrona já está ocupada.\n")
		}

		jan[n-1] = 1
		fmt.Print("Poltrona marcada!")
		time.Sleep(1 * time.Second)
		clearTerminal()

	case 2: 
		var n int 
		
		fmt.Print("\nDigite o número da poltrona na qual você deseja sentar.\n")
		fmt.Scan(&n)

		if n < 1 || n > 24 {
			fmt.Println("Esta poltrona não existe, escolha outra entre 1 e 24.")
			continue
		}

		if meio[n-1] == 1 {
			fmt.Print("Esta poltrona já está ocupada.\n")
		}

		meio[n-1] = 1
		fmt.Print("Poltrona marcada.")
		time.Sleep(1 * time.Second)
		clearTerminal()
	}
	completo(jan)
	completo(meio)
	}
}