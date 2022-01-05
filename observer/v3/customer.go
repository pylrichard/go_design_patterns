package main

import "fmt"

type Customer struct {
	Id string
}

func (c *Customer) Update(name string) {
	fmt.Printf("Sending email to Customer %s for Item %s\n", c.Id, name)
}

func (c *Customer) GetId() string {
	return c.Id
}