package event

import (
	"errors"
	"fmt"

	bool2 "github.com/sikasjc/utility/bool"
)

type Manager struct {
	id          string
	isRunning   *bool2.AtomicBool
	ch          chan Event
	callbackMap map[uint]func(event Event)
	cap         int
}

func NewManager(id string, size int) *Manager {
	m := &Manager{
		id:          id,
		ch:          make(chan Event, size),
		callbackMap: make(map[uint]func(event Event)),
		cap:         size,
	}
	return m
}

func (m *Manager) Num() int {
	return len(m.ch)
}

func (m *Manager) Cap() int {
	return cap(m.ch)
}

func (m *Manager) Register(eventType uint, callback func(Event)) error {
	if callback == nil {
		return fmt.Errorf("register eventType %d fail, callback is nil", eventType)
	}
	if _, ok := m.callbackMap[eventType]; ok {
		return fmt.Errorf("register eventType %d fail, already registered", eventType)
	}
	m.callbackMap[eventType] = callback
	return nil
}

func (m *Manager) UnRegister(eventType uint) {
	delete(m.callbackMap, eventType)
}

func (m *Manager) RecvNonBlock(event Event) error {
	if m.isRunning.IsUnSet() {
		return errors.New("manager is not running")
	}
	select {
	case m.ch <- event:
		return nil
	default:
	}
	return errors.New("event chan is full")
}

func (m *Manager) RecvBlock(event Event) error {
	if m.isRunning.IsUnSet() {
		return errors.New("manager is not running")
	}
	m.ch <- event
	return nil
}

func (m *Manager) Chan() chan Event {
	return m.ch
}

func (m *Manager) Do(event Event) {
	if event == nil {
		return
	}
	if callback, ok := m.callbackMap[event.Type()]; ok {
		callback(event)
	}
}

func (m *Manager) IsRunning() bool {
	return m.isRunning.IsSet()
}

func (m *Manager) Set(b bool) {
	m.isRunning.SetTo(b)
}
