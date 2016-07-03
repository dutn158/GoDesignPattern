package main

import (
	"testing"
)

func TestComposite(t *testing.T) {
	var student3 = &Student{name: "Student3", score: 11}
	var student4 = &Student{name: "Student4", score: 9}

	var class2 = &Class{}

	class2.students = append(class2.students, *student3)
	class2.students = append(class2.students, *student4)

	if class2.calculateScore() != 19 {
		t.Error("Expect 19 but receive ", class2.calculateScore())
	}

}
