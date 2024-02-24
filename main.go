package main

import "fmt"

import (
	_ "github.com/DataDog/agent-payload/v5/gogen"
	_ "github.com/mheffner/go-simple-metrics"
)

func main() {
	fmt.Println("hello world")
}
