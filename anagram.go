// Author Randy Coburn
// Date: 24/09/2016
// Description: This is a simple program that will
// 				check if 2 strings are a anagram of
//				each other.
// Example:	anagram randycoburn coburnrandy
// http://www.dictionary.com/browse/anagram

package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	// collect the words from the arguments
	// Lower case everything to rule out caps.
	word1 := strings.ToLower(os.Args[1])
	word2 := strings.ToLower(os.Args[2])

	// run a sorting function that will sort the letters
	if sorter(word1) == sorter(word2) {
		// Print True and exit good.
		fmt.Println("True")
		os.Exit(0)
	} else {
		// Print False and exit bad
		fmt.Println("False")
		os.Exit(1)
	}

}

func sorter(word string) string {
	// Take the word and split it to bytes.
	split := strings.Split(word,"")
	// sorts a slice of strings in increasing order.
	// https://golang.org/pkg/sort/#Strings
	sort.Strings(split)
	// Join then up again and send back to the caller.
	return strings.Join(split,"")
}