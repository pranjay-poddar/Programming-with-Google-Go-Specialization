// Write two goroutines which have a race condition when executed concurrently. Explain what the race condition is and how it can occur.

// Explanation
// ================
// The program spawns two routines that each increment the global Counter variable twice. When both routines are done running,
// the program displays the value of the global Counter variable. I have added a billionth of a second pause into the loop.
// I put the pause right after the routine reads the global Counter variable and stores a local copy. This pause in the loop
// has caused the program to fail.

package main

import (
	"fmt"
	"sync"
	"time"
)

//Wait ... the wait group
var Wait sync.WaitGroup

//Counter ... a global counter
var Counter int

func main() {

	for routine := 1; routine <= 2; routine++ {

		Wait.Add(1)
		go Routine(routine)
	}

	Wait.Wait()
	fmt.Printf("Final Counter: %d\n", Counter)
}

//Routine ... a routine
func Routine(id int) {

	for count := 0; count < 2; count++ {

		value := Counter
		time.Sleep(1 * time.Nanosecond)
		value++
		Counter = value
	}

	Wait.Done()
}
