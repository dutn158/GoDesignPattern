package main

import (
	"fmt"
)

type Observer interface {
	Update(i int64)
}

type Subject interface {
	AddObserver(o Observer)

	RemoveObserver(o Observer)

	NotifyObserver(i int64)
}

type MainActivity struct {
	observers []Observer
}

type Fragment struct {
}

func (main *MainActivity) AddObserver(o Observer) {
	main.observers = append(main.observers, o)
}

func (main *MainActivity) RemoveObserver(o Observer) {
	// main.observers = remove(main.observers, o)
}

func (main *MainActivity) NotifyObserver(i int64) {
	for _, r := range main.observers {
		r.Update(i)
	}
}

func (fragment *Fragment) Update(i int64) {
	fmt.Printf("%v", i)
}

func main() {
	s := &MainActivity{}

	o1 := &Fragment{}
	o2 := &Fragment{}

	s.AddObserver(o1)
	s.AddObserver(o2)

	s.NotifyObserver(15)

	s.NotifyObserver(2016)

}
