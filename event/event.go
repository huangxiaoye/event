package event

import (
	"fmt"
	"reflect"
	"sync"
	"log"
	"errors"
)

type Handler func (...interface{}) bool

type Eventer interface {
	AddEvent(name string, handler Handler) (error)
	FireEvent(name string, params ...interface{})
	ClearEvent(name string) error
	Die()
}

type Event struct {
	eventsMutex sync.Mutex
	Handlers map[string][]Handler
}

func NewEvent() *Event{
	return &Event{Handlers:make(map[string][]Handler)}
}

func (e *Event) Die() {
	e.Handlers = make(map[string][]Handler)
}

func (e *Event) ClearEvent (name string) error {
	if _, ok := e.Handlers[name]; ok {
		e.eventsMutex.Lock()
		delete(e.Handlers, name)
		e.eventsMutex.Unlock()
		return nil
	}

	log.Fatalln("Not found in this object")
	return errors.New("Not found Handlers!")
}

func (e *Event) FireEvent(name string, params ...interface{}) {
	if h, ok := e.Handlers[name]; ok {
		for _, handle := range h {
			if flag := handle(params...); flag {
				break
			}
		}
		return
	}

	log.Fatalln(fmt.Sprintf("Can not found %s event!", name))
	return
	
}

func (e *Event) AddEvent(name string, handler Handler) error {
	validateHandler(handler)
	if h, ok := e.Handlers[name]; ok {
		e.eventsMutex.Lock()
		e.Handlers[name] = append(h, handler)
		e.eventsMutex.Unlock()
		return nil
	}
	e.eventsMutex.Lock()
	e.Handlers[name] = append(e.Handlers[name], handler)
	e.eventsMutex.Unlock()
	return nil
}

func validateHandler(h Handler) {
	if reflect.TypeOf(h).Kind() != reflect.Func {
		panic("handler must be a callable func")
	}
}