package engine

import (
	"bufio"
	"github.com/stretchr/testify/assert"
	"os"
	"strconv"
	"testing"
)

func TestLoop(t *testing.T) {
	printCmd1 := PrintCommand("first print")
	printCmd2 := PrintCommand("second print")
	addCmd := AddCommand{
		A: 10,
		B: 12,
	}

	loop := new(Loop)
	loop.Start()
	assert.Equal(t, false, loop.stop)
	assert.Equal(t, 0, len(loop.q.commands))

	loop.Post(addCmd)
	loop.Post(printCmd1)
	loop.Post(printCmd2)

	assert.Equal(t, 3, len(loop.q.commands))
	loop.AwaitFinish()
	assert.Equal(t, true, loop.stop)
	assert.Equal(t, 0, len(loop.q.commands))

	var outputLines []string
	readFile, _ := os.Open("results.txt")
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		outputLines = append(outputLines, fileScanner.Text())
	}

	assert.Equal(t, string(printCmd1), outputLines[0])
	assert.Equal(t, string(printCmd2), outputLines[1])
	assert.Equal(t, strconv.Itoa(addCmd.A+addCmd.B), outputLines[2])
}
