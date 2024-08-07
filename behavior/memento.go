package behavior

import "fmt"

// memento or history

type Originator struct {
	state string
}

func (e *Originator) createMemento() *Memento {
	return &Memento{state: e.state}
}

func (e *Originator) restoreMemento(m *Memento) {
	e.state = m.getSavedState()
}

func (e *Originator) setState(state string) {
	e.state = state
}

func (e *Originator) getState() string {
	return e.state
}

type Memento struct {
	state string
}

func (m *Memento) getSavedState() string {
	return m.state
}

type Caretaker struct {
	mementoArray []*Memento
}

func (c *Caretaker) addMemento(m *Memento) {
	c.mementoArray = append(c.mementoArray, m)
}

func (c *Caretaker) getMemento(index int) *Memento {
	return c.mementoArray[index]
}

func DoMemento() {
	careTaker := &Caretaker{mementoArray: make([]*Memento, 0)}
	originator := &Originator{state: "A"}

	fmt.Printf("Originator Current state: %s\n", originator.getState())
	careTaker.addMemento(originator.createMemento())

	originator.setState("B")
	fmt.Printf("Originator Current state: %s\n", originator.getState())
	careTaker.addMemento(originator.createMemento())

	originator.setState("C")
	fmt.Printf("Originator Current state: %s\n", originator.getState())
	careTaker.addMemento(originator.createMemento())

	originator.restoreMemento(careTaker.getMemento(1))
	fmt.Printf("Originator Current state: %s\n", originator.getState())

	originator.restoreMemento(careTaker.getMemento(0))
	fmt.Printf("Originator Current state: %s\n", originator.getState())
}
