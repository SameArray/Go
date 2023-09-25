package main

import (
	"fmt"
)

type Operation struct {
	Action string // "load", "store", "delete", "range"
	Key    interface{}
	Value  interface{}
}

type Result struct {
	Value interface{}
	OK    bool
}

func main() {
	users := make(map[interface{}]interface{})

	opChannel := make(chan Operation)
	resChannel := make(chan Result)

	go func() {
		for op := range opChannel {
			switch op.Action {
			case "load":
				value, ok := users[op.Key]
				resChannel <- Result{Value: value, OK: ok}
			case "store":
				users[op.Key] = op.Value
				resChannel <- Result{OK: true}
			case "delete":
				delete(users, op.Key)
				resChannel <- Result{OK: true}
			case "range":
				for key, value := range users {
					fmt.Printf("User ID: %v, Name: %v\n", key, value)
				}
				resChannel <- Result{OK: true}
			}
		}
	}()

	opChannel <- Operation{Action: "store", Key: 1, Value: "Alice"}
	<-resChannel
	opChannel <- Operation{Action: "store", Key: 2, Value: "Bob"}
	<-resChannel
	opChannel <- Operation{Action: "store", Key: 3, Value: "Charlie"}
	<-resChannel

	opChannel <- Operation{Action: "load", Key: 2}
	result := <-resChannel
	if result.OK {
		fmt.Println("User with ID 2:", result.Value)
	}

	opChannel <- Operation{Action: "delete", Key: 1}
	<-resChannel

	opChannel <- Operation{Action: "load", Key: 1}
	result = <-resChannel
	if !result.OK {
		fmt.Println("User with ID 1 not found.")
	}

	opChannel <- Operation{Action: "range"}
	<-resChannel
	close(opChannel)
}
