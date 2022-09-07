/*
The function call io.Copy(dst, src) read from src and write to dst. Use it instead of ioutil.ReadAll to copy the response body to os.Stdout
*/

package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
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
		fmt.Printf("%d", w)
	}
}
