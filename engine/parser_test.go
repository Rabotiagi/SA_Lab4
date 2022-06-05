package engine

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	res := Parse("print hello!")
	assert.Equal(t, PrintCommand("hello!"), res)

	res = Parse("printt hello!")
	assert.Equal(t, PrintCommand("SYNTAX ERROR: Unknown instruction"), res)

	res = Parse("add 1 8")
	assert.Equal(t, AddCommand{A: 1, B: 8}, res)

	res = Parse("add a 8")
	assert.Equal(t, PrintCommand("SYNTAX ERROR: strconv.Atoi: parsing \"a\": invalid syntax"), res)

	res = Parse("add 8 b")
	assert.Equal(t, PrintCommand("SYNTAX ERROR: strconv.Atoi: parsing \"b\": invalid syntax"), res)

	res = Parse("add 9")
	assert.Equal(t, PrintCommand("SYNTAX ERROR: Not enough arguments"), res)

	res = Parse("")
	assert.Equal(t, PrintCommand("SYNTAX ERROR: Not enough arguments"), res)
}
