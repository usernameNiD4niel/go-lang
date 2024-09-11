package main

import (
	"fmt"

	"github.com/usernameNiD4niel/go-lang/greetings"
)

func main() {
	message := greetings.Hello("Daniel")
	fmt.Println(message)
}
