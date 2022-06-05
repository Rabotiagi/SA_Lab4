package main

import (
	"bufio"
	"eventloop/engine"
	"os"
	"strconv"
	"strings"
)

func parse(commandLine string) engine.Command {
	parts := strings.Fields(commandLine)

	switch parts[0] {
	case "print":
		return engine.PrintCommand(parts[1])
	case "add":
		num1, err := strconv.Atoi(parts[1])
		num2, err := strconv.Atoi(parts[2])
		if err != nil {
			return engine.PrintCommand("SYNTAX ERROR: " + err.Error())
		}
		return engine.AddCommand{A: num1, B: num2}
	}

	return engine.PrintCommand("SYNTAX ERROR: Unknown instruction")
}

func main() {
	loop := new(engine.Loop)
	loop.Start()

	if input, err := os.Open("test.txt"); err == nil {
		defer input.Close()
		scanner := bufio.NewScanner(input)
		for scanner.Scan() {
			commandLine := scanner.Text()
			cmd := parse(commandLine)
			loop.Post(cmd)
		}
	}

	loop.AwaitFinish()
}
