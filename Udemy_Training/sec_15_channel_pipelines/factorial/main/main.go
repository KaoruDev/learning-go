package main

func main() {

}

func factorial(num int) <- chan int {
    result := make(chan int)

    go func() {
        total := 1

        for i := num; i > 0; i-- {
            total *= i
        }

        result <- total
        close(result)
    }()

    return result
}

func factorialChan(num int, queue chan int) int {
    if (num <= 1) {
        close(queue)
        return num
    }

    go func() {
        queue <- num
    }()

    return factorialChan(num - 1, queue)
}
