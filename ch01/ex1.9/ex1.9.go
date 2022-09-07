/*
Modify to print HTTP status code
*/

package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {

		if !strings.HasPrefix(url, "http://") {
			url = "http://" + url

		}

		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stdout, "fetch: %v\n", err)
			os.Exit(1)
		}
		//b, err := ioutil.ReadAll(resp.Body)
		w, err := io.Copy(os.Stdout, resp.Body)
		if err != nil {
			fmt.Fprintf(os.Stdout, "Fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("byte written: %d\n", w)
		fmt.Printf("Status code: %d\n", resp.StatusCode)
	}
}
