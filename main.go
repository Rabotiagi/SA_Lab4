package main

import (
	"eventloop/engine"
)

func main() {
	loop := new(engine.Loop)

	loop.Start()

	loop.Post(engine.PrintCommand("hello"))
	loop.Post(&engine.AddCommand{A: 4, B: 5})
	loop.Post(engine.PrintCommand("hello2"))

	loop.AwaitFinish()
}
