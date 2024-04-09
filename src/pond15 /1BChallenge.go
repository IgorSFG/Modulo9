package main

import (
    "fmt"
    "os"
    "sync"
)

const (
  numLines int = 1000000000
  goDiv int = 100
  goLast int = goDiv - 1
)

func main() {
    file, err := os.Create("1BChallenge.txt")
    if err != nil {
        fmt.Println("Erro ao criar arquivo:", err)
        return
    }
    defer file.Close()

    var wg sync.WaitGroup
    linesPerRoutine := numLines / goDiv
    remainder := numLines % goDiv

    for i := 0; i < goDiv; i++ {
        wg.Add(1)
        start := i * linesPerRoutine
        end := start + linesPerRoutine
        if i == numLast {
            end += remainder
        }
        go func(start, end int) {
            defer wg.Done()
            for j := start; j < end; j++ {
                fmt.Fprintf(file, "Linha %d\n", j+1)
            }
        }(start, end)
    }

    wg.Wait()
    fmt.Println("Desafio de 1 Bilhão de linhas concluído.")
}
