package main

import "fmt"

func greet(name string) string {
	return "Hi, " + name + " po hahhaha"
}

func main() {
	fmt.Println("Hello World!")
	fmt.Println(greet("daniel"))
	var num uint8 = 255
	fmt.Println(num)
}
