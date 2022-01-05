package main

type Observer interface {
	Update(string)
	GetId() string
}