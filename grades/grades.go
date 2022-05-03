package grades

import (
	"fmt"
	"sync"
)

type Student struct {
	Id        int
	FirstName string
	LastName  string
	Grades    []Grade
}

type GradeType string

type Grade struct {
	Title string
	Type  GradeType
	Score float32
}

const GradeQuiz = GradeType("quiz")
const GradeTest = GradeType("Test")
const GradeExam = GradeType("Exam")

func (s Student) Average() float32 {
	var result float32
	for _, grade := range s.Grades {
		result += grade.Score
	}

	return result / float32(len(s.Grades))
}

type Students []Student

var (
	students      Students
	studentsMutex sync.Mutex
)

func (s Students) GetById(id int) (*Student, error) {
	for i := range s {
		if s[i].Id == id {
			return &s[i], nil
		}
	}
	return nil, fmt.Errorf("student with ID %d not found", id)
}
