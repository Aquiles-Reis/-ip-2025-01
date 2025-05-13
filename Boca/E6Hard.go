package main 

import "fmt"

func ordenar(vet[]int, n int) {
for j := 0; j < n; j++ {
for i := 0; i < n - 1 - j; i++ {
	if vet[i] > vet[i+1] {
		vet[i], vet[i+1] = vet[i+1], vet[i]
		}
	}
}

}


func main (){
	
	var (
		n int
	)
	fmt.Scan(&n)
	vet := make ([]int, n)
	for i := 0; i < n; i ++ {
		fmt.Scan(&vet[i])
	}
	ordenar(vet, n)

	for _, x := range vet {
		fmt.Printf("%d ", x)
	}

}