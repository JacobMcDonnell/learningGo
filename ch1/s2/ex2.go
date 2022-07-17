// Prints cmd input with index value
package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + strconv.Itoa(i) + ": " + os.Args[i]
		sep = "\n"
	}
	fmt.Println(s)
}
