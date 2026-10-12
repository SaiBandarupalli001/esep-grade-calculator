package esepunittests

import "testing"

func TestGetGradeA(t *testing.T) {
	expected_value := "A"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 100, Assignment)
	gradeCalculator.AddGrade("exam 1", 100, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 100, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeB(t *testing.T) {
	expected_value := "B"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 80, Assignment)
	gradeCalculator.AddGrade("exam 1", 81, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 85, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeF(t *testing.T) {
	expected_value := "F"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 50, Assignment)
	gradeCalculator.AddGrade("exam 1", 55, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 59, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGradeBoundaries(t *testing.T) {
	tests := []struct {
		score    int
		expected string
	}{
		{90, "A"},
		{89, "B"},
		{80, "B"},
		{79, "C"},
		{70, "C"},
		{69, "D"},
		{60, "D"},
		{59, "F"},
	}

	for _, test := range tests {
		calculator := NewGradeCalculator()
		calculator.AddGrade("assignment", test.score, Assignment)
		calculator.AddGrade("exam", test.score, Exam)
		calculator.AddGrade("essay", test.score, Essay)

		result := calculator.GetFinalGrade()

		if result != test.expected {
			t.Errorf("For score %d, expected %s but got %s",
				test.score, test.expected, result)
		}
	}
}

func TestGradeTypeString(t *testing.T) {
	tests := []struct {
		gradeType GradeType
		expected  string
	}{
		{Assignment, "assignment"},
		{Exam, "exam"},
		{Essay, "essay"},
	}

	for _, test := range tests {
		result := test.gradeType.String()

		if result != test.expected {
			t.Errorf("Expected %s but got %s", test.expected, result)
		}
	}
}

func TestPassFailGrading(t *testing.T) {
	tests := []struct {
		score    int
		expected string
	}{
		{100, "Pass"},
		{70, "Pass"},
		{69, "Fail"},
		{0, "Fail"},
	}

	for _, test := range tests {
		calculator := NewGradeCalculator(true)

		calculator.AddGrade("assignment", test.score, Assignment)
		calculator.AddGrade("exam", test.score, Exam)
		calculator.AddGrade("essay", test.score, Essay)

		result := calculator.GetFinalGrade()

		if result != test.expected {
			t.Errorf("For score %d, expected %s but got %s",
				test.score, test.expected, result)
		}
	}
}
