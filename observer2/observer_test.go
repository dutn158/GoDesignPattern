package main

import (
	"testing"
)

func TestObserver(t *testing.T) {
	s := &MainActivity{}

	o1 := &Fragment{}
	o2 := &Fragment{}

	s.AddObserver(o1)
	s.AddObserver(o2)

	s.NotifyObserver(15)

	if len(s.observers) != 2 {
		t.Error("Expect 2, but receive ", len(s.observers))
	}

}
