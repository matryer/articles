package main

import "sync"

type thing struct {
	m sync.Mutex
}

func main() {
}

func (t *thing) doSomething() {
	defer un(lock(&t.m))

}

func lock(m *sync.Mutex) *sync.Mutex {
	m.Lock()
	return m
}
func un(m *sync.Mutex) {
	m.Unlock()
}
