package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Reading in Full Sentences
	// reader := bufio.NewReader(os.Stdin)
	// fmt.Println("Simple Shell")

	// for {
	// 	fmt.Print("-> ")
	// 	text, _ := reader.ReadString('\n')
	// 	// convert CRLF to LF
	// 	text = strings.Replace(text, "\n", "", -1)

	// 	if strings.Compare("hi", text) == 0 {
	// 		fmt.Println("hello, Yourself")
	// 	}
	// }

	// Reading Single UTF-8 Encoded Unicode Characters
	// reader := bufio.NewReader(os.Stdin)
	// char, _, err := reader.ReadRune()
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// // print out the unicode value i.e. A -> 65, a -> 97
	// fmt.Println(char)

	// switch char {
	// case 'A':
	// 	fmt.Println("A Key Pressed")
	// 	break
	// case 'a':
	// 	fmt.Println("a Key Pressed")
	// 	break
	// }

	// Using Bufio's Scanner
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
