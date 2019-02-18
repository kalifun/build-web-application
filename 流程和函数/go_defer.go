package main

import (
	"fmt"
	"os"
)

func main() bool {
	file.Open("file")
	defer file.Close()
	if failurex {
		return false
	}
	if failurey {
		return false
	}
	return true
}