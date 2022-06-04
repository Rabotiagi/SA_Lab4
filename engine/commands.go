package engine

import (
	"fmt"
	"strconv"
)

type Command interface {
	Execute(h Handler)
}

type PrintCommand string
type AddCommand struct {
	A, B int
}
type stopCommand struct{}

func (pc PrintCommand) Execute(h Handler) {
	fmt.Println(string(pc))
}

func (ac AddCommand) Execute(h Handler) {
	res := ac.A + ac.B
	h.Post(PrintCommand(strconv.Itoa(res)))
}

func (sc stopCommand) Execute(h Handler) {
	h.(*Loop).stop = true
}
