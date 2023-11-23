package main

import (
	"flag"
	"fmt"
)

type JobType string

func (j *JobType) Set(s string) error {
	switch s {
	case "kafkaConsumer":
		*j = KafkaConsumer
	case "subscriptions":
		*j = Subscriptions
	}
	return nil
}

// implemented some methods in flag.Value interface
func (j *JobType) String() string {

	return string(*j)
}

const (
	KafkaConsumer JobType = "kafkaConsumer"
	Subscriptions JobType = "subscriptions"
)

func main() {

	var jobRunner JobType
	flag.Var(&jobRunner, "jobRunner", "Job to run. Valid Job:kafkaConsumer or subscriptions")

	flag.Parse()
	fmt.Printf("%v", string(jobRunner) == "")

}
