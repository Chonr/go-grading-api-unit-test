package grade

type MockRepository struct{}

func (m *MockRepository) GetGradeByStudentID(studentID string) (*Response, error) {
	return &Response{
		StudentID: studentID,
		Total:     85,
		Grade:     "A",
	}, nil
}

func (m *MockRepository) InsertGrade(g Response, homework, midterm, final float64) error {

	return nil
}
