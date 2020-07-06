package main

import (
	"fmt"
	"os"
)

// file samples
func main() {
	// os file
	file, _ := os.Open("/Users/niaoshuai")
	defer file.Close()

}
