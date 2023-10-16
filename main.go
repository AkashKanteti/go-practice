package main

import (
	"fmt"
	"time"
)

func main() {
	date := time.Now()
	fmt.Printf("%s\n", date)
	date = date.AddDate(0, 1, -date.Day())
	var dates string = date.Format("02-01-2006")
	fmt.Printf("%s", dates)
}
