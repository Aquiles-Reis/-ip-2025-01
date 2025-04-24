package main

import "fmt"

// Função para realizar a busca binária
func buscaBinaria(array []int, alvo int) int {
    esquerda := 0
    direita := len(array) - 1

    for esquerda <= direita {
        meio := esquerda + (direita-esquerda)/2
        if array[meio] == alvo {
            return meio // Retorna o índice do elemento encontrado
        } else if array[meio] < alvo {
            esquerda = meio + 1
        } else {
            direita = meio - 1
        }
    }
    return -1 // Retorna -1 se o elemento não for encontrado
}

func main() {
    array := []int{1, 3, 5, 7, 9, 11}
    alvo := 5

    resultado := buscaBinaria(array, alvo)
    if resultado != -1 {
        fmt.Printf("Elemento %d encontrado no índice %d.\n", alvo, resultado)
    } else {
        fmt.Printf("Elemento %d não encontrado.\n", alvo)
    }
}