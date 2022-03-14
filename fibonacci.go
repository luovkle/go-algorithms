package main

import "fmt"

func fibonacci(n int) []int {
    var serie []int
    a, b := 0, 1
    for b < n {
        serie = append(serie, b)
        a, b = b, a + b
    }
    return serie
}

func main() {
    serie := fibonacci(100)
    fmt.Println(serie)
}
