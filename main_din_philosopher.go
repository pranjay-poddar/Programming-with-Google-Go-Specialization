package main

import (
	"fmt"
	"sync"
)

type Chops struct{ sync.Mutex }
type Philo struct {
	name            int
	leftCS, rightCS *Chops
}

func (p Philo) eat(wg *sync.WaitGroup, ch chan int) {
	for i := 0; i < 3; i++ {
		ch <- p.name
		p.leftCS.Lock()
		p.rightCS.Lock()
		fmt.Printf("ID %d finishing eating: %d time\n", p.name, i+1)
		p.rightCS.Unlock()
		p.leftCS.Unlock()
		<-ch
	}
	wg.Done()
}

func main() {

	CSticks := make([]*Chops, 5)
	for i := 0; i < 5; i++ {
		CSticks[i] = new(Chops)
	}

	philos := make([]*Philo, 5)
	for i := 0; i < 5; i++ {
		philos[i] = &Philo{i + 1, CSticks[i], CSticks[(i+1)%5]}
	}

	/*
		Channel with buffer:2, doesn't allow more than 2 philosophers
		to eat concurrently.
	*/
	currently_eating := make(chan int, 2)
	var wg sync.WaitGroup
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go philos[i].eat(&wg, currently_eating)
	}
	wg.Wait()
}
