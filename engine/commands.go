package engine

import (
	"fmt"
	"strconv"
)

type Command interface {
	Execute(h Handler)
}

type printCommand string
type addCommand struct {
	a, b int
}

func (pc printCommand) Execute(h Handler) {
	fmt.Println(string(pc))
}

func (ac addCommand) Execute(h Handler) {
	res := ac.a + ac.b
	h.Post(printCommand(strconv.Itoa(res)))
}
