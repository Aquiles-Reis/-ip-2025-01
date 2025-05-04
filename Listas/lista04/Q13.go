package main

import (
    "fmt"
    "sort"
)

type Empregado struct {
    numero int
    meses  int
}

func main() {
    var empregados []Empregado

    for {
        var numero, meses int
        fmt.Scan(&numero, &meses)

        if numero == 0 && meses == 0 {
            break
        }

        empregados = append(empregados, Empregado{numero, meses})

        if len(empregados) >= 100 {
            break
        }
    }

    // Ordenar por número de meses de trabalho (do menor para o maior)
    sort.Slice(empregados, func(i, j int) bool {
        return empregados[i].meses < empregados[j].meses
    })

    // Garantir que imprime no máximo os três empregados mais recentes
    fmt.Println("Os três empregados mais recentes:")
    for i := 0; i < 3 && i < len(empregados); i++ {
        fmt.Println("Empregado:", empregados[i].numero, "- Meses de trabalho:", empregados[i].meses)
    }
}