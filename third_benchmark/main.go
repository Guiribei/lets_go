package main

import (
    "fmt"
    "time"
)

func main() {
    start := time.Now()

    counter := 0
    for i := 0; i < 10000; i++ {
        fmt.Printf("Hello #%v\n", counter)
        counter++
    }

    elapsed := time.Since(start)
    fmt.Printf("Sequential execution took %s\n", elapsed)
}

