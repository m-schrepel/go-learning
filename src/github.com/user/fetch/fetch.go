// Fetch prints the content found at a URL.
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	for i, url := range os.Args[1:] {
		// Prepend http to the url if it isn't already there.
		// No clever way to deal with http vs https at the moment
		// at least not for this example
		if strings.HasPrefix(os.Args[i], "http") != true {
			url = "http://" + url
		}

		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		b, err := io.Copy(os.Stdout, resp.Body)
		s := resp.StatusCode
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%d", s)
		fmt.Printf("%v", b)
	}
}
