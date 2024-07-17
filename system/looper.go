package system

import (
	"time"
)

type Looper struct {
	frequency      time.Duration
	timeAccumulate time.Duration
	do             func()
}

func (l *Looper) SetFrequency(frequency time.Duration) {
	l.frequency = frequency
}

func (l *Looper) Update() {
	l.timeAccumulate += DeltaTime
	l.doLoop()
}

func (l *Looper) doLoop() {
	if l.frequency < 0 {
		return
	}
	if l.timeAccumulate >= l.frequency {
		l.timeAccumulate -= l.frequency
		l.do()
		l.doLoop()
	}
}

func NewLooper(frequency time.Duration, do func()) *Looper {
	return &Looper{
		do:             do,
		frequency:      frequency,
		timeAccumulate: 0,
	}
}
