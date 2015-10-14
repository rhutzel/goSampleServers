package main

import (
	"fmt"
	"net/rpc"
)

type Obj struct { Value int }

func main() {
	client, _ := rpc.Dial("tcp", ":1234")
	var result Obj
	client.Call("GoCounter.StepIncrement", &Obj{2}, &result)
	fmt.Printf("%d\n", result.Value)
}
