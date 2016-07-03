package observer

import (
	"math/rand"
)

type observer interface {
	update() int
}

type number interface {
	getNumber() int
}

type DigitObserver struct {
	generator number
}

type numberGenerator struct {
	observers []observer
}

func (self *numberGenerator) AddObserver(observer observer) {
	self.observers = append(self.observers, observer)
}

func (self *numberGenerator) notifyObservers() []int {
	var result []int
	for _, observer := range self.observers {
		result = append(result, observer.update())
	}
	return result
}

type randomNumberGenerator struct {
	*numberGenerator
}

func NewRandomNumberGenerator() *randomNumberGenerator {
	return &randomNumberGenerator{&numberGenerator{}}
}

func (self *randomNumberGenerator) getNumber() int {
	return rand.Intn(50)
}

func (self *randomNumberGenerator) Execute() []int {
	return self.notifyObservers()
}

func (self *DigitObserver) update() int {
	return self.generator.getNumber()
}
