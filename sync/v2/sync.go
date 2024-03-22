package sync

import "sync"

type Counter struct {
	mux   sync.Mutex
	value int
}

func NewCounter() *Counter {
	return &Counter{}
}

func (c *Counter) Inc() {
	c.mux.Lock()
	defer c.mux.Unlock()
	c.value++
}

func (c *Counter) Value() int {
	return c.value
}
