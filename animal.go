/*Write a program which allows the user to get information about a predefined set of animals. Three animals are 
predefined, cow, bird, and snake. Each animal can eat, move, and speak. The user can issue a request to find out
one of three things about an animal: 1) the food that it eats, 2) its method of locomotion, and 3) the sound it 
makes when it speaks. The following table contains the three animals and their associated data which should be 
hard-coded into your program.

Your program should present the user with a prompt, “>”, to indicate that the user can type a request. Your 
program accepts one request at a time from the user, prints out the answer to the request, and prints out a 
new prompt. Your program should continue in this loop forever. Every request from the user must be a single 
line containing 2 strings. The first string is the name of an animal, either “cow”, “bird”, or “snake”. The 
second string is the name of the information requested about the animal, either “eat”, “move”, or “speak”. 
Your program should process each request by printing out the requested data.

You will need a data structure to hold the information about each animal. Make a type called Animal which is
a struct containing three fields:food, locomotion, and noise, all of which are strings. Make three methods 
called Eat(), Move(), and Speak(). The receiver type of all of your methods should be your Animal type. The
Eat() method should print the animal’s food, the Move() method should print the animal’s locomotion, and 
the Speak() method should print the animal’s spoken sound. Your program should call the appropriate method
when the user makes a request.
*/

package main

import(
	"fmt"
	"strings"
)

type Animal struct{
	food string
    locomotion string
	noise string
}

func (animal Animal) Eat(){
	fmt.Println(animal.food)
} 


func (animal Animal) Speak(){
	fmt.Println(animal.noise)
}


func (animal Animal) Move(){
	fmt.Println(animal.locomotion)
} 

func main(){
	var start,animal_name,animal_info string
	var animal Animal
	for{
		fmt.Print("Please enter '>' to continue or any other key to exit: ")
		fmt.Scan(&start)
		start = strings.ToLower(start)
		if(start != ">"){
			fmt.Println("Program exited!")
				return	
			}
		fmt.Print("Please enter animal name:")
		fmt.Scan(&animal_name)
		animal_name = strings.ToLower(animal_name)

		fmt.Print("Please enter the info: ")
		fmt.Scan(&animal_info)
	    animal_info = strings.ToLower(animal_info)
	    switch animal_name{
		case "cow":
			animal = Animal{"grass", "walk", "moo"}
		case "bird":
			animal = Animal{"worms", "fly", "peep"}
		case "snake":
			animal = Animal{"mice", "slither", "hsss"}
		default:
			fmt.Printf("%s is not in (cow, bird, snake)", animal_name)
		return
		}	
		
		switch animal_info{
		case "eat":
			animal.Eat()
		case "speak":
			animal.Speak()
		case "move":
			animal.Move()
		default:
			fmt.Println("%s is not in (Eat, Speak, Move)",animal_info)
		return
		}
	}
}
