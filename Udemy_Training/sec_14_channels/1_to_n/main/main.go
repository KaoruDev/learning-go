package main

import (
    "fmt"
    "time"
    "math/rand"
)

/**
This shows how we could have multiple consumers to a channel
 */

func main() {
    totalConsumers := 10
    queue := make(chan int)
    semaphore := make(chan bool)

    go func() {
        for i := 0; i < 10000; i++ {
            queue <- i
        }
        close(queue)
    }()

    for consumer := 0; consumer < totalConsumers; consumer++ {
        go func() {
            for work := range queue {
                time.Sleep(time.Microsecond * time.Duration(rand.Intn(100)))
                fmt.Println(work)
            }
            semaphore <- true
        }()
    }

    for done := 0; done < totalConsumers; done++ {
       <- semaphore
    }
}
