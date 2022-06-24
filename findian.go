package main

import (
  "bufio"
    "os"
	"fmt"
	"strings"
	)

	func main(){
	    scanner := bufio.NewScanner(os.Stdin)
		fmt.Print("Please enter the string: ")    
		scanner.Scan()
        line := strings.ToLower(scanner.Text())
		
		str := []rune(line)
		found := 0
		size := len(str)
        if str[0] != 'i'{
		    println("Not Found!")
		    return 
		}
		if str[size-1]!='n'{
			println("Not Found!")
			return
		}
		for i:= 1;i < size-1 ;i++{
		    if str[i] == 'a'{
		    found = 1
		   }
		}	   
		if found == 0{
	    fmt.Println("Not Found!")
		}else{
			fmt.Println("Found!")
		}
	}