/*
Modify the echo program to also print os.Args[0], the name of the command that invoke it.
*/
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	//changge os.Args[1:] to os.Args
	fmt.Println(strings.Join(os.Args, " "))
}
