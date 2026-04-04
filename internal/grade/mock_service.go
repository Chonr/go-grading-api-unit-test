package grade

import (
	"errors"
)

type MockService struct {
	ShouldReturnError bool
}

func (m *MockService) CheckGrade(studentID string) (*Response, error) {

	if m.ShouldReturnError {
		return nil, errors.New("not found")
	}
	return &Response{
		StudentID: studentID,
		Total:     90,
		Grade:     "A",
	}, nil
}

func (m *MockService) SubmitGrade(req Request) (*Response, error) {

	if m.ShouldReturnError {
		return nil, errors.New("not found")
	}
	return &Response{
		StudentID: req.StudentID,
		Total:     80,
		Grade:     "A",
	}, nil
}
