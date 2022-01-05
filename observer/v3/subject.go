package main

type Subject interface {
	Register(o Observer)
	Cancel(o Observer)
	NotifyAll()
}