package main

import (
    "sync"
    "fmt"
)

func main() {
    c := make(chan int)

    wg := sync.WaitGroup{}

    wg.Add(1)
    go func() {
        for i := 0; i < 10; i++ {
            c <- i
        }
        wg.Done()
    }()

    wg.Add(1)
    go func() {
        for i := 0; i < 10; i++ {
            c <- i
        }
        wg.Done()
    }()

    go func() {
        wg.Wait()
        close(c)
    }()

    for n := range(c) {
        fmt.Println(n)
    }
}
