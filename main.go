package main

import (
	"fmt"
)

type JobType string

// implemented some methods in flag.Value interface

const (
	KafkaConsumer JobType = "kafkaConsumer"
	Subscriptions JobType = "subscriptions"
)

func main() {
	// fmt.Printf("%v", jobRunner)
	// ok := funs()
	// fmt.Printf("%v", ok)
	switch funs() {
	case KafkaConsumer:
		fmt.Printf(string(KafkaConsumer))
	case Subscriptions:
		fmt.Printf(string(Subscriptions))
	}
}

func funs() JobType {
	return "subscriptions"
}
