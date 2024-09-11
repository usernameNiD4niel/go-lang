package golang

import "fmt"

func Hello(mess string) string {
	message := fmt.Sprintf("Hello %v welcome to the family", mess)
	return message
}
