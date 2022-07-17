package main

import (
	"fmt"
	"bufio"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
)

func main () {
	fmt.Print("Write a series of integers, split by space > ")
	
	read := bufio.NewReader(os.Stdin)
	input, _ := read.ReadString('\n')

	s := strings.Split(strings.TrimSpace(input), " ")
	sli := make([]int, 0, len(s))

	for _, str := range s {
		num, _ := strconv.Atoi(str)
		sli = append(sli, num)
	}

	const par = 4
	n := int(math.Max(math.Ceil(float64(len(s))/float64(par)), 1))

	// WaitGroup to synchronize all go routines
	var wg sync.WaitGroup

	//Partition the given array to approximately equal size
	for i := 0; i < par; i++ {
		from := int(math.Min(float64(i*n), float64(len(sli))))
		to := int(math.Min(float64((i+1)*n), float64(len(sli))))

		//Add one go routine to WaitGroup wg.
		wg.Add(1)

		go func(arr []int) {
			fmt.Println("Will sort: ", arr)
			sort.Ints(arr)

			//Removes go routine from WaitGroup wg.
			wg.Done()
		}(sli[from:to])
	}

	// Wait for all go routines to finish
	wg.Wait()

	sort.Ints(sli)
	fmt.Println("Sorted: ", sli)
}