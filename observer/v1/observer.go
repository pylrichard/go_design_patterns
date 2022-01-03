package v1

import "fmt"

type ISubject interface {
	Register(o IObserver)
	Remove(o IObserver)
	Notify(msg string)
}

type IObserver interface {
	Update(msg string)
}

type Subject struct {
	observers []IObserver
}

func (s *Subject) Register(o IObserver) {
	s.observers = append(s.observers, o)
}

func (s *Subject) Remove(o IObserver) {
	for i, ob := range s.observers {
		if ob == o {
			s.observers = append(s.observers[:i], s.observers[i + 1:]...)
		}
	}
}

func (s *Subject) Notify(msg string) {
	for _, o := range s.observers {
		o.Update(msg)
	}
}

type Observer1 struct {}

func (Observer1) Update(msg string) {
	fmt.Printf("Observer1: %s\n", msg)
}

type Observer2 struct {}

func (Observer2) Update(msg string) {
	fmt.Printf("Observer2: %s\n", msg)
}