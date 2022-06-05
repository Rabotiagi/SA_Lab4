package engine

import (
	"strconv"
	"strings"
)

func Parse(commandLine string) Command {
	parts := strings.Fields(commandLine)

	if len(parts) < 2 {
		return PrintCommand("SYNTAX ERROR: Not enough arguments")
	}

	switch parts[0] {
	case "print":
		return PrintCommand(parts[1])
	case "add":
		if len(parts) < 3 {
			return PrintCommand("SYNTAX ERROR: Not enough arguments")
		}
		num1, err := strconv.Atoi(parts[1])
		if err != nil {
			return PrintCommand("SYNTAX ERROR: " + err.Error())
		}
		num2, err := strconv.Atoi(parts[2])
		if err != nil {
			return PrintCommand("SYNTAX ERROR: " + err.Error())
		}
		return AddCommand{A: num1, B: num2}
	}

	return PrintCommand("SYNTAX ERROR: Unknown instruction")
}
