package main

import "fmt"

func PrintHelloWorld() string {
	return "Hello, world.\n"
}

func main() {
	fmt.Printf(PrintHelloWorld)
}
