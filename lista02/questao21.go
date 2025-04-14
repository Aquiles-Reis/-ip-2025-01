package main 

import "fmt"

func main (){

	medias := []float64{4.0, 6.0, 7.5, 9.0, 10.0}
	conceitos := []string{"Reprovado E", "Reprovado D", "Aprovado C", "Aprovado B", "Aprovado A"}
	
	var ident int
	var n1, n2, n3, ex float64
	
	fmt.Scan(&ident, &n1, &n2, &n3, &ex)

	media := (n1 + (n2 * 2) + (n3 * 3) + ex) / 7

		conceito := "Indefinido"
		for i, limite := range medias {
			if media <= limite {
				conceito = conceitos [i] 
				break 
			}
		} 
		if conceito == "Indefinido" {
			conceito = conceitos [len(conceitos)-1]
		}


		fmt.Printf("%d %.2f %.2f %.2f %.2f %.2f %s", ident, n1, n2, n3, ex, media, conceito)
}