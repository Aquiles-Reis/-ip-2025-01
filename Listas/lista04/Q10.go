package main 

import "fmt"

func Fibonacci (n int) []int {
	fibSeq := make([]int, n)
	fibSeq [0], fibSeq [1] = 1, 1

	for i:=2; i < n; i++ {
		fibSeq[i] = fibSeq[i - 1] + fibSeq [i - 2] 
	}
	return fibSeq
}

func main (){
	numTermos := 50
	sequencia := Fibonacci(numTermos)
	fmt.Println(sequencia)

}