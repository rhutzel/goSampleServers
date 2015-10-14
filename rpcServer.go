package main

import (
	"net"
	"net/rpc"
)

type Obj struct { Value int }

type Counter struct {
	count int
}

func (c *Counter) StepIncrement(input *Obj, output *Obj) error {
	c.count += input.Value
	output.Value = c.count
	return nil
}

func main() {
	server := rpc.NewServer()
	server.RegisterName("GoCounter", new(Counter))
	listener, _ := net.Listen("tcp", "127.0.0.1:1234")
	server.Accept(listener)
}
