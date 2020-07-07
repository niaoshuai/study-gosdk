package os

import (
	"os"
	"testing"
)

// file samples
func TestFile(t *testing.T) {
	// os file
	file, _ := os.Open("/Users/niaoshuai")
	defer file.Close()

}
