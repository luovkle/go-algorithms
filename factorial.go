package main

import "fmt"

func factorial(n uint) uint {
    if n <= 1 {
        return 1
    }
    return n * factorial(n - 1)
}

func main() {
    var n uint
    fmt.Printf("Number: ")
    _, err := fmt.Scanf("%d", &n)
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Printf("Factorial: %d\n", factorial(n))
    }
}
