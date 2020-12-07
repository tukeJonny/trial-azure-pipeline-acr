package main

import "fmt"

func main() {
	var input string
	for {
		fmt.Printf("Please input> ")
		fmt.Scanln(&input)
		fmt.Println(input)
	}
}
