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