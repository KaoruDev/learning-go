package main

import "fmt"

func main() {
    multiChanelWay()
    singleChannelWay()
}

func singleChannelWay() {
    numChan := incrementor()
    sum := sum(numChan)
    fmt.Printf("Single chan way, sum is %d\n", sum)
}

func sum(ints <-chan int) int {
    result := 0

    for num := range ints {
        result += num
    }

    return result
}

func multiChanelWay() {
    numChan := incrementor()
    sumChan := sumChan(numChan)

    for sum := range sumChan {
        fmt.Printf("Multi channel way, sum is: %d\n", sum)
    }
}

func sumChan(ints <-chan int) <-chan int {
	sumChan := make(chan int)

	go func() {
        sum := 0

        for i := range ints {
            sum += i
        }

        sumChan <- sum
        close(sumChan)
    }()

	return sumChan
}

func incrementor() <-chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			out <- i
		}
		close(out)
	}()
	return out
}
