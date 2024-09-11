package main

import "fmt"

func Hello(name string) string {
	message := fmt.Sprintf("Hello, %v nice meeting you!", name)
	return message
}
