package main

import "runtime"

type Counter interface {
	Add(uint64)
	Read() uint64
	Init()
}

type ChannelCounter struct {
	ch     chan func()
	number uint64
}

func (c *ChannelCounter) Init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	panic("implement me")
}

func NewChannelCounter() Counter {
	counter := &ChannelCounter{make(chan func(), 100), 0}
	go func(counter *ChannelCounter) {
		for f := range counter.ch {
			f()
		}
	}(counter)
	return counter
}

func (c *ChannelCounter) Add(num uint64) {
	c.ch <- func() {
		c.number = c.number + num
	}
}

func (c *ChannelCounter) Read() uint64 {
	ret := make(chan uint64)
	c.ch <- func() {
		ret <- c.number
		close(ret)
	}
	return <-ret
}

func main() {

}
