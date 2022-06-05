package main

import (
	"bufio"
	"eventloop/engine"
	"os"
)

func main() {
	loop := new(engine.Loop)
	loop.Start()

	if input, err := os.Open("test.txt"); err == nil {
		defer input.Close()
		scanner := bufio.NewScanner(input)
		for scanner.Scan() {
			commandLine := scanner.Text()
			cmd := engine.Parse(commandLine)
			loop.Post(cmd)
		}
	}

	loop.AwaitFinish()
}
