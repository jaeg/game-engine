package main

import (
	"fmt"

	"github.com/jaeg/game-engine/event"
)

const (
	Test event.EventType = iota
)

var m *event.QueuedEventManager

type TestEventData struct {
	x int
	y int
}

func (t TestEventData) GetType() event.EventType {
	return Test
}

type TestListener struct {
}

func (t *TestListener) HandleEvent(data event.EventData) error {
	fmt.Println("Handling event: ", data)
	m.QueueEvent(data)
	return nil
}

func main() {
	m = &event.QueuedEventManager{}

	testListener := &TestListener{}

	m.RegisterListener(testListener, Test)
	m.RegisterListener(testListener, Test)
	m.RegisterListener(testListener, Test)

	//Test event send
	testEventData := TestEventData{x: 5, y: 6}
	m.QueueEvent(testEventData)
	m.HandleQueue()
	m.HandleQueue()
	//m.UnregisterListener(testListener, Test)
	m.UnregisterListenerFromAll(testListener)
}
