package main

import (
	"fmt"
	"sync"
)

type ChopS struct{ sync.Mutex }

type Philo struct {
	id              int
	leftCS, rightCS *ChopS
	counter         int
}

var wg sync.WaitGroup

func main() {
	CSticks := make([]*ChopS, 5)
	for i := 0; i < 5; i++ {
		CSticks[i] = new(ChopS)
	}
	philos := make([]*Philo, 5)
	for i := 0; i < 5; i++ {
		philos[i] = &Philo{i, CSticks[i], CSticks[(i+1)%5], 1}
	}

	wg.Add(15)

	for _, p := range philos {
		go eat(*p)
	}

	wg.Wait()
}

func eat(p Philo) {
	for {
		if p.counter <= 3 {
			p.leftCS.Lock()
			p.rightCS.Lock()

			fmt.Printf("starting to eat %d \n", p.id)
			fmt.Printf("finishing eating %d \n", p.id)

			p.rightCS.Unlock()
			p.leftCS.Unlock()
			p.counter = p.counter + 1
			wg.Done()
		} else {
			return
		}
	}
}
