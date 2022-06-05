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

	// ONLY FOR RUNNING TESTS
	//f, err := os.OpenFile("results.txt", os.O_APPEND|os.O_WRONLY, 0644)
	//if err != nil {
	//	f, _ = os.Create("results.txt")
	//	f.WriteString(string(pc))
	//} else {
	//	f.WriteString("\n" + string(pc))
	//}
	//
	//f.Close()
}

func (ac AddCommand) Execute(h Handler) {
	res := ac.A + ac.B
	h.Post(PrintCommand(strconv.Itoa(res)))
}

func (sc stopCommand) Execute(h Handler) {
	h.(*Loop).stop = true
}
