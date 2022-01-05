package main

import "fmt"

type Item struct {
	Observers []Observer
	Name      string
	InStock   bool
}

func NewItem(name string) *Item {
	return &Item {
		Name: name,
	}
}

func (i *Item) UpdateAvailability() {
	fmt.Printf("Item %s is now in Stock\n", i.Name)
	i.InStock = true
	i.NotifyAll()
}

func (i *Item) Register(o Observer) {
	i.Observers = append(i.Observers, o)
}

func (i *Item) Cancel(o Observer) {
	i.Observers = remove(i.Observers, o)
}

func (i *Item) NotifyAll() {
	for _, observer := range i.Observers {
		observer.Update(i.Name)
	}
}

func remove(observers []Observer, o Observer) []Observer {
	length := len(observers)
	for i, observer := range observers {
		if o.GetId() == observer.GetId() {
			observers[length - 1], observers[i] = observers[i], observers[length - 1]
			return observers[:length - 1]
		}
	}

	return observers
}