package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"
)

func main() {
	// define our initial map where the words are the key and the count is the value
	counts := make(map[string]int)
	// for each file input as the args on the command line...
	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}
		for _, line := range strings.Split(string(data), " ") {
			counts[line]++
		}
	}
	// Now we attempt a sort. First we must turn values into a slice
	// since maps are randomized on purpose (can't guarantee order)
	keys := make([]int, len(counts))

	// Build out the slice with the values
	for _, value := range counts {
		keys = append(keys, value)
	}

	// sort in descending order
	sort.Sort(sort.Reverse(sort.IntSlice(keys)))

	// get the reversed map
	var reversedCounts = reverse(counts)

	// for all the keys in the slice / map print out the key and value
	for _, n := range keys {
		if n > 1 {
			fmt.Println("Key:", n, "Value:", reversedCounts[n])
		}
	}
}

// reverse the key/value pairs so we can list the lines by the key of the string
// which is the number of occurrences
func reverse(m map[string]int) map[int]string {
	reversedCounts := make(map[int]string)
	for k, v := range m {
		reversedCounts[v] = k
	}
	return reversedCounts
}
