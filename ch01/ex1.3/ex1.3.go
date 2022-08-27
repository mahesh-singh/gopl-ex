/*
Experiment to measure the difference in runing time in inefficient version vs string.join
*/

package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	s := ""
	beforeTime := time.Now()
	for _, arg := range os.Args {
		s = s + arg + " "
	}
	fmt.Println(s)
	totalTime := time.Now().Sub(beforeTime)
	fmt.Println(totalTime)

	beforeTime = time.Now()
	fmt.Println(strings.Join(os.Args, " "))
	totalTime = time.Now().Sub(beforeTime)
	fmt.Println(totalTime)
}
