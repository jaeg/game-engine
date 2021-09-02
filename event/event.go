package event

import (
	"fmt"
)

type EventType int

// EventData - interface representing event data.  Each data struct needs to return its type as an EventType (int)
type EventData interface {
	GetType() EventType
}

// EventListener - interface representing an event listener.  Each event listener needs a HandleEvent function
type EventListener interface {
	HandleEvent(EventData) error
}

// EventManager - Entry point for registering event listeners and sending events
type EventManager struct {
	listeners map[EventType][]EventListener // Key is the event
}

// RegisterListener - registers a struct that can be represented by the EventListener
// interface to the specified event type
func (m *EventManager) RegisterListener(listener EventListener, eventType EventType) {
	if m.listeners == nil {
		m.listeners = make(map[EventType][]EventListener)
	}

	m.listeners[eventType] = append(m.listeners[eventType], listener)
}

// SendEvent - Sends the event data to the event listeners registered to its type.
func (m *EventManager) SendEvent(data EventData) {
	if m.listeners == nil {
		m.listeners = make(map[EventType][]EventListener)
		return
	}

	for _, v := range m.listeners[data.GetType()] {
		err := v.HandleEvent(data)
		if err != nil {
			fmt.Println(err)
		}
	}
}

// UnregisterListener - unregisters a struct that can be represented by the EventListener
// interface from the specified event type's events
func (m *EventManager) UnregisterListener(listener EventListener, eventType EventType) {
	if m.listeners == nil {
		m.listeners = make(map[EventType][]EventListener)
		return
	}

	for k := 0; k < len(m.listeners[eventType]); k++ {
		if m.listeners[eventType][k] == listener {
			m.listeners[eventType] = append(m.listeners[eventType][:k], m.listeners[eventType][k+1:]...)
			k--
		}
	}
}

// UnregisterListenerFromAll - unregisters a struct that can be represented by the EventListener
// interface from all events
func (m *EventManager) UnregisterListenerFromAll(listener EventListener) {
	if m.listeners == nil {
		m.listeners = make(map[EventType][]EventListener)
		return
	}

	for eventType := range m.listeners {
		for k := 0; k < len(m.listeners[eventType]); k++ {
			if m.listeners[eventType][k] == listener {
				m.listeners[eventType] = append(m.listeners[eventType][:k], m.listeners[eventType][k+1:]...)
				k--
			}
		}
	}

}
