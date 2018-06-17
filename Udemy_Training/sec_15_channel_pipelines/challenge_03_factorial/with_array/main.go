package with_array

import (
"sync"
"fmt"
)

// Runs with array answer
func Run(workCount int) {
    in := gen(workCount)

    f := factorial(in)

    for n := range f {
        fmt.Println(n)
    }
}

func gen(workCount int) <-chan int {
    out := make(chan int)
    go func() {
        for i := 0; i < workCount; i++ {
            for j := 3; j < 13; j++ {
                out <- j
            }
        }
        close(out)
    }()
    return out
}

func factorial(in <-chan int) <-chan int {
    out := make(chan int)

    var calculations = []<-chan int{}

    // fan out
    for n := range in {
        calculations = append(calculations, calculate(n))
    }

    var wg sync.WaitGroup

    // fan in
    for _, calc := range calculations {
        wg.Add(1)

        go func(myCalc <- chan int) {
            for result := range myCalc {
                out <- result
                // wg.Done() // I had this, which won't work because there is no guarantee that myCalc only has 1 value
            }
            wg.Done()
        }(calc)
    }

    go func() {
        wg.Wait()
        close(out)
    }()

    return out
}

func calculate(number int) <- chan int {
    out := make(chan int)

    go func() {
        out <- fact(number)
        close(out)
    }()

    return out
}

func fact(n int) int {
    total := 1
    for i := n; i > 0; i-- {
        total *= i
    }
    return total
}

/*
CHALLENGE #1:
-- Change the code above to execute 1,000 factorial computations concurrently and in parallel.
-- Use the "fan out / fan in" pattern

CHALLENGE #2:
WATCH MY SOLUTION BEFORE DOING THIS CHALLENGE #2
-- While running the factorial computations, try to find how much of your resources are being used.
-- Post the percentage of your resources being used to this discussion: https://goo.gl/BxKnOL
*/
