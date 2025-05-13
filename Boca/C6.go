package main

import "fmt"

func main (){

	var alunos, aprovados, exame, reprovados int
	var n1, n2 float64

	fmt.Scan(&alunos)
	
	for i:=1; i <= alunos; i++ {
		fmt.Scan(&n1, &n2)

		media := (n1 + n2) / 2
		
		fmt.Printf("Aluno %d: ", i)
		if media <= 3 {
			fmt.Print("Reprovado \n")
			reprovados++
		} else if media > 3 && media < 7 {
			fmt.Print("Exame \n")
			exame++
		} else {
			fmt.Print("Aprovado \n")
			aprovados++
		}
	} 
	fmt.Printf("Total Aprovados: %d \n", aprovados)
	fmt.Printf("Total Exame: %d \n", exame)
	fmt.Printf("Total Reprovados: %d \n", reprovados)
}