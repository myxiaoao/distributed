package grades

func init() {
	students = []Student{
		{
			Id:        1,
			FirstName: "Nick",
			LastName:  "Carter",
			Grades: []Grade{
				{
					Title: "Quiz 1",
					Type:  GradeQuiz,
					Score: 85,
				}, {
					Title: "Final Exam",
					Type:  GradeExam,
					Score: 95,
				}, {
					Title: "Quiz 2",
					Type:  GradeQuiz,
					Score: 82,
				},
			},
		},
		{
			Id:        2,
			FirstName: "Cooper",
			LastName:  "Make",
			Grades: []Grade{
				{
					Title: "Quiz 1",
					Type:  GradeQuiz,
					Score: 85,
				}, {
					Title: "Final Exam",
					Type:  GradeExam,
					Score: 95,
				}, {
					Title: "Test",
					Type:  GradeTest,
					Score: 68,
				},
			},
		},
	}
}
