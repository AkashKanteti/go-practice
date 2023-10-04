package main

import (
	"bufio"
	"fmt"
	"os"

	"rsc.io/quote"
)

func main() {
	fmt.Printf(quote.Go())
	scanner := bufio.NewScanner(os.Stdin)

	// Write the data to standard input
	scanner.Scan()
	text := scanner.Text()
	fmt.Println(text)
	// scanner.Scan()

}

//echo 33 | ./main current usage
