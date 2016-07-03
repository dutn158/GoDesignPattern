package main

import (
	"fmt"
)

type composite interface {
	calculateScore() float64
}

type Student struct {
	name  string
	score float64
}

type Class struct {
	students []Student
}

func (student *Student) calculateScore() float64 {
	return student.score
}

func (class *Class) calculateScore() float64 {
	totalScore := 0.0
	for _, v := range class.students {
		totalScore += v.score
	}

	return totalScore
}

func main() {

	var student1 = &Student{name: "Student1", score: 7}
	var student2 = &Student{name: "Student2", score: 9}

	var class1 = &Class{}
	class1.students = append(class1.students, *student1)
	class1.students = append(class1.students, *student2)

	fmt.Println(student1.calculateScore())
	fmt.Println(student2.calculateScore())
	fmt.Println(class1.calculateScore())

}
