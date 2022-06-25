/*Write a program which prompts the user to first enter a name, and then enter an address. Your program should create
 a map and add the name and address to the map using the keys “name” and “address”, respectively. Your program should
  use Marshal() to create a JSON object from the map, and then your program should print the JSON object.*/

  package main

  import (
	  "fmt"
	  "os"
	  "encoding/json"
	  "bufio"
	  "strings"
  )

  func main(){
      scanner := bufio.NewScanner(os.Stdin)
	  var lines []string
	  fmt.Print("Please enter a name: ") 
	  for i:=1;i<3;i++ { 
	  				scanner.Scan()
	  			    line := strings.ToLower(scanner.Text())
					if len(line) == 0 {
						break
		            }
		  if i ==1{
			  fmt.Print("Please enter an address: ")  
		  } 
			        lines = append(lines, line)
		  }
	  
      m := make(map[string]string) 
	  m[lines[0]] = lines[1]
	  fmt.Println("Map: ",m)
	  fmt.Print("JSON: ")
	  jsonStr, err := json.Marshal(m)
	  if err==nil{
	  fmt.Println(string(jsonStr))
	}
  }