/*Write a program which prompts the user to enter a string. The program searches through the entered string for the characters
 ‘i’, ‘a’, and ‘n’. The program should print “Found!” if the entered string starts with the character ‘i’, ends with the
 character ‘n’, and contains the character ‘a’. The program should print “Not Found!” otherwise. The program should not be
 case-sensitive, so it does not matter if the characters are upper-case or lower-case.

Examples: The program should print “Found!” for the following example entered strings, “ian”, “Ian”, “iuiygaygn”, “I d skd a
 efju N”. The program should print “Not Found!” for the following strings, “ihhhhhn”, “ina”, “xian”.
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Please enter the string: ")
	scanner.Scan()
	line := strings.ToLower(scanner.Text())

	str := []rune(line)
	found := 0
	size := len(str)
	if str[0] != 'i' {
		println("Not Found!")
		return
	}
	if str[size-1] != 'n' {
		println("Not Found!")
		return
	}
	for i := 1; i < size-1; i++ {
		if str[i] == 'a' {
			found = 1
		}
	}
	if found == 0 {
		fmt.Println("Not Found!")
	} else {
		fmt.Println("Found!")
	}
}
