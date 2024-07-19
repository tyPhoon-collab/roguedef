package system

import (
	"time"
)

type Looper struct {
	Frequency      time.Duration
	timeAccumulate time.Duration
	do             func()
}

func (l *Looper) Update() {
	l.timeAccumulate += ScaledDeltaTime()
	l.doLoop()
}

func (l *Looper) doLoop() {
	if l.Frequency < 0 {
		return
	}
	if l.timeAccumulate >= l.Frequency {
		l.timeAccumulate -= l.Frequency
		l.do()
		l.doLoop()
	}
}

func (l *Looper) WithFrequency(frequency time.Duration) *Looper {
	l.Frequency = frequency
	return l
}

func NewLooper(frequency time.Duration, do func()) *Looper {
	return &Looper{
		do:             do,
		Frequency:      frequency,
		timeAccumulate: 0,
	}
}
