package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"sync"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Swap(sli []int, i int) {
	tmp := sli[i]
	sli[i] = sli[i+1]
	sli[i+1] = tmp
}

func BubbleSort(sli []int, wg *sync.WaitGroup) {
	length := len(sli)
	for i := 0; i < length-1; i++ {
		for j := 0; j < length-i-1; j++ {
			if sli[j] > sli[j+1] {
				Swap(sli, j)
			}
		}
	}
	if wg != nil {
		wg.Done()
	}
}

func main() {
	//a := []int{1, 4, 2, 3, 7, 3, 2, 3, 4, 21, 3, 4, 5, 8, 9, 3, 2, 1}
	////a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18}
	//fmt.Println(a)
	fmt.Print("Enter integers (seperated by comma): ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	inputStr := scanner.Text()
	inputs := strings.Split(inputStr, ",")
	var a []int
	for _, v := range inputs {
		value, err := strconv.Atoi(v)
		if err != nil {
			continue
		}
		a = append(a, value)
	}
	n := len(a)
	partSize := int(math.Ceil(float64(n) / 4))
	//fmt.Println(partSize)
	wg := sync.WaitGroup{}
	for i := 0; i < 4; i++ {
		sli := a[i*partSize : min((i+1)*partSize, n)]
		fmt.Printf("Subarray[%d]: ", i)
		fmt.Println(sli)
		if len(sli) > 0 {
			wg.Add(1)
			go BubbleSort(sli, &wg)
		}
	}
	wg.Wait()
	BubbleSort(a, nil)
	fmt.Println("Sorted array:", a)
}
