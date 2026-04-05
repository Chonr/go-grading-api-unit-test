package grade

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateGrade_TableDriven(t *testing.T) {
	tests := []struct {
		name     string
		homework float64
		midterm  float64
		final    float64
		expected string
	}{
		{"Grade A", 80, 70, 90, "A"},
		{"Grade B", 75, 75, 75, "B"},
		{"Grade C", 65, 65, 65, "C"},
		{"Grade D", 55, 55, 55, "D"},
		{"Grade F", 40, 40, 40, "F"},
		{"Invalid Under", -1, 50, 50, "Invalid"},
		{"Invalid Over", 101, 50, 50, "Invalid"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, grade := CalculateGrade(tt.homework, tt.midterm, tt.final)
			assert.Equal(t, tt.expected, grade)
		})
	}
}

func TestCheckGrade(t *testing.T) {
	mockRepo := &MockRepository{}
	service := NewGradeService(mockRepo)

	res, err := service.CheckGrade("6609650269")

	assert.NoError(t, err)
	assert.Equal(t, "6609650269", res.StudentID)
	assert.Equal(t, "A", res.Grade)
}

func TestCheckGrade_EmptyID(t *testing.T) {
	service := NewGradeService(&MockRepository{})
	_, err := service.CheckGrade("")
	assert.Error(t, err)
	assert.Equal(t, "student ID is required", err.Error())
}

type MockRepositoryError struct{ MockRepository }

func (m *MockRepositoryError) GetGradeByStudentID(id string) (*Response, error) {
	return nil, errors.New("db error")
}

func TestCheckGrade_RepoError(t *testing.T) {
	service := NewGradeService(&MockRepositoryError{})
	_, err := service.CheckGrade("6609650269")
	assert.Error(t, err)
}

func (m *MockRepositoryError) InsertGrade(g Response, h, mid, f float64) error {
	return errors.New("insert failed")
}

func TestSubmitGrade_RepoError(t *testing.T) {
	service := NewGradeService(&MockRepositoryError{})

	req := Request{StudentID: "6609650269", Homework: 80, Midterm: 70, Final: 90}
	res, err := service.SubmitGrade(req)

	assert.Error(t, err)
	assert.Nil(t, res)
	assert.Equal(t, "insert failed", err.Error())
}

func TestSubmitGrade_Success(t *testing.T) {
	mockRepo := &MockRepository{}
	service := NewGradeService(mockRepo)

	req := Request{
		StudentID: "6609650269",
		Homework:  80,
		Midterm:   70,
		Final:     90,
	}

	res, err := service.SubmitGrade(req)

	assert.NoError(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, "A", res.Grade)
}

type MockRepositoryInsertError struct{ MockRepository }

func (m *MockRepositoryInsertError) InsertGrade(g Response, h, mi, f float64) error {
	return errors.New("insert failed")
}

func TestSubmitGrade_InsertError(t *testing.T) {
	mockRepo := &MockRepositoryInsertError{}
	service := NewGradeService(mockRepo)

	req := Request{StudentID: "6609650269", Homework: 80, Midterm: 70, Final: 90}
	res, err := service.SubmitGrade(req)

	assert.Error(t, err)
	assert.Nil(t, res)
	assert.Equal(t, "insert failed", err.Error())
}
