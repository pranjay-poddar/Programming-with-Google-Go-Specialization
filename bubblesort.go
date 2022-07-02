/*Write a Bubble Sort program in Go. The program
should prompt the user to type in a sequence of up to 10 integers. The program
should print the integers out on one line, in sorted order, from least to
greatest. Use your favorite search tool to find a description of how the bubble
sort algorithm works.

As part of this program, you should write a
function called BubbleSort() which
takes a slice of integers as an argument and returns nothing. The BubbleSort() function should modify the slice so that the elements are in sorted
order.

A recurring operation in the bubble sort algorithm is
the Swap operation which swaps the position of two adjacent elements in the
slice. You should write a Swap() function which performs this operation. Your Swap()
function should take two arguments, a slice of integers and an index value i which
indicates a position in the slice. The Swap() function should return nothing, but it should swap
the contents of the slice in position i with the contents in position i+1.*/

package main

import (
	"fmt"
)

func Swap(num1 *int,num2 *int){
var x int = *num1
*num1 = *num2
*num2 = x
}

func BubbleSort(sli []int){
for i:= 0;i < len(sli)-1;i++{
	for j:= 0;j < len(sli)-i-1;j++{
         if sli[j] > sli[j+1]{
           Swap(&sli[j],&sli[j+1])
		 }
	}
}
}

func main(){
	
	fmt.Println("Please enter 10 integers: ")
	var arr[10] int
	for i := 0;i < 10;i++{
		fmt.Scan(&arr[i])
	}
	a := arr[:]
    BubbleSort(a)
	fmt.Print("Sorted array: ")
	for i := 0;i < 10;i++{
		fmt.Print(arr[i] , " ")
	}
	fmt.Print("\n")
	
}