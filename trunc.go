package main
import "fmt"

func main(){
	var num float64;
	fmt.Println("Please enter a number: ")
	fmt.Scanf("%f",&num)
	fmt.Println("The turncated value is:",int(num))
}