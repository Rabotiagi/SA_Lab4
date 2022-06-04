package engine

import "sync"

type CommandQueue struct {
	me             sync.Mutex
	commands       []Command
	emptyFlag      bool
	notEmptySignal chan struct{}
}

func (cq *CommandQueue) pull() Command {
	cq.me.Lock()
	defer cq.me.Unlock()

	if len(cq.commands) == 0 {
		cq.emptyFlag = true
		cq.me.Unlock()

		<-cq.notEmptySignal
		cq.me.Lock()
	}

	res := cq.commands[0]
	cq.commands[0] = nil
	cq.commands = cq.commands[1:]
	return res
}

func (cq *CommandQueue) push(c Command) {
	cq.me.Lock()
	defer cq.me.Unlock()
	cq.commands = append(cq.commands, c)

	if cq.emptyFlag {
		cq.notEmptySignal <- struct{}{}
	}
}

func (cq *CommandQueue) empty() bool {
	cq.me.Lock()
	defer cq.me.Unlock()
	return len(cq.commands) == 0
}
