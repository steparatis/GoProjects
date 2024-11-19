package main

import "fmt"

type Checker struct {
	items []Checkable
}

type Measurable interface {
	GetMetricks() string
}

type Checkable interface {
	Measurable
	Ping() error
	GetId() string
	Health() bool
}

func (c *Checker) Add(checkInt Checkable) {
	c.items = append(c.items, checkInt)
}

func (c Checker) Stringer() string{
	var str string

	for _, item := range c.items {
		str += item.GetId()
	}

	return str

}

func (c Checker) Check() {
	for _, item := range c.items {
		if item.Health() == false {
			fmt.Println(item.GetId() + " не работает")
		}
	}
}