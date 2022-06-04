package engine

type Handler interface {
	Post(c Command)
}

type Loop struct {
	q *CommandQueue

	stop       bool
	stopSignal chan struct{}
}

func (l *Loop) Start() {
	l.q = &CommandQueue{
		notEmptySignal: make(chan struct{}),
	}
	l.stopSignal = make(chan struct{})
	go func() {
		for !l.stop || !l.q.empty() {
			cmd := l.q.pull()
			cmd.Execute(l)
		}
		l.stopSignal <- struct{}{}
	}()
}

func (l *Loop) Post(cmd Command) {
	l.q.push(cmd)
}

func (l *Loop) AwaitFinish() {
	l.Post(stopCommand{})
	l.stop = true
	<-l.stopSignal
}
