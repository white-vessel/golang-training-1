package main

import (
	"bufio"
	"fmt"
	"os"
)

func CountBooks(books []string) map[string]int {

	bookCount := make(map[string]int)

	defer fmt.Println("Counting complete!")

	// count the occurrences of each title.
	for _, book := range books {
		bookCount[book]++
	}

	return bookCount
}

func main() {

	books := []string{}

	// use `for scanner.Scan()` to keep reading
	for {
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		line := scanner.Text()

		//type 0 to end the receiving loop
		if line == "0" {
			break
		}

		//handling the case where the input is empty
		if line == "" {
			continue
		}

		books = append(books, line)

	}

	counts := CountBooks(books)

	fmt.Println(counts)

}
