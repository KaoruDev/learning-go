package main

import "fmt"

func main() {
	correctWay()

	// fmt.Println("========== WRONG WAY TWO")
	// wrongWayTwo()

	// fmt.Println("========== WRONG WAY THREE")
	// wrongWayThree()

    fmt.Println("========== WRONG WAY THREE FIXED")
    wrongThreeFixed()
}

func correctWay() {
	workChan := make(chan int)
	semaphore := make(chan bool)

	go func() {
		for i := 0; i < 10; i++ {
			workChan <- i
		}
		semaphore <- true
	}()

	go func() {
		for i := 0; i < 10; i++ {
			workChan <- i
		}
		semaphore <- true
	}()

	go func() {
		<-semaphore
		<-semaphore
		close(workChan)
	}()

	for n := range workChan {
		fmt.Println(n)
	}
}

func wrongWay() {
	workChan := make(chan int)
	semaphore := make(chan bool)

	go func() {
		for i := 0; i < 10; i++ {
			workChan <- i
		}
		semaphore <- true
	}()

	go func() {
		for i := 0; i < 10; i++ {
			workChan <- i
		}
		semaphore <- true
	}()

	/*
		Because workChan is a unbuffered channel, it blocks threads from putting on more than 1 item on its queue.
		Thus we lock everything up because the threads above won't be able to finish their iterations, and therefore won't
		be able to put anything on the semaphore channel
	*/
	<-semaphore
	<-semaphore
	close(workChan)

	for n := range workChan {
		fmt.Println(n)
	}
}

func wrongWayTwo() {
	workChan := make(chan int)
	semaphore := make(chan bool)

	go func() {
		for i := 0; i < 10; i++ {
			workChan <- i
		}
		semaphore <- true
	}()

	go func() {
		for i := 0; i < 10; i++ {
			workChan <- i
		}
		semaphore <- true
	}()

	// This won't work either because the for loop won't exit until the workChan is closed.
	for n := range workChan {
		fmt.Println(n)
	}

	<-semaphore
	<-semaphore
	close(workChan)
}

func wrongWayThree() {
	someChan := make(chan int)

	/*
	   this throws an error... I'm not entirely sure why, perhaps you can't
	   add to a channel on the same go routine as you pull it out

	   Okay so the reason why this throws an error is because someChan <- 1 will actually block
	   until something pulls the value off the channel. This is probably because the channel is unbuffered
	*/

	someChan <- 1
	close(someChan)

	for i := range someChan {
		fmt.Println(i)
	}
}

/*
By making the channel a "buffered" channel, it fixes the problem because now the thread that puts thing on the channel
is not longer blocked.
 */

func wrongThreeFixed() {
    someChan := make(chan int, 1)

    someChan <- 1
    // Uncommenting the lines below exceeds the buffer thus we get into another deadlock situation
    // someChan <- 1
    // someChan <- 1
    // someChan <- 1
    close(someChan)

    for i := range someChan {
        fmt.Println(i)
    }
}
