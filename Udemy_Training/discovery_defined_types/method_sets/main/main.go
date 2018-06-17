package main

import (
	"fmt"
)

type Mutex struct {
	locked bool
}

func (m *Mutex) plock() bool {
	if m.locked {
		fmt.Println("mutex is already locked")
		return false
	}

	m.locked = true
	fmt.Println("mutex is now locked")
	return true
}

func (m *Mutex) unlock() {
	fmt.Println("Unlocked mutex!")
	m.locked = false
}

func (m *Mutex) lock() {
	pM := &m
	fmt.Println(pM);
	fmt.Println(*pM);
	fmt.Println(*m);
}

type NewMutex Mutex

type PrintableMutex struct {
	Mutex
}

func main() {
	mutex := Mutex{false}
	printableMutex := PrintableMutex{mutex}

	fmt.Println("====== Locking printable locks")
	printableMutex.plock()
	printableMutex.Mutex.plock()
	printableMutex.plock()

	fmt.Println("\n\n")

	fmt.Println("====== Pointers")
	printPointer(&mutex)
	mutex.lock()
	printPointer(&printableMutex)


	fmt.Println("\n\n")

	fmt.Println("====== locking og mutex")
	mutex.plock()
	fmt.Println(&mutex)

	newM := NewMutex(mutex)
	fmt.Println("NewMutex only has access to locked: ", newM.locked);
}

func printPointer(p interface{}) {
	fmt.Println(&p)
}
