
# ESEP Grade Calculator

A simple grade calculator library written in Go. It calculates a student's final letter grade using assignments, exams, and essays.

## Requirements

- Go 1.24.4 or later
- Git for version control

## Project Files

- `README.MD`: Project documentation
- `grade-calculator/go.mod`: Go module configuration
- `grade-calculator/grade_calculator.go`: Grade calculator implementation
- `grade-calculator/grade_calculator_test.go`: Unit tests

## Grade Calculation

The calculator is intended to calculate the average for each grade category and combine them using these weights:

- Assignments: 50%
- Exams: 35%
- Essays: 15%

The resulting numerical grade is converted using this scale:

- A: 90 or above
- B: 80 to below 90
- C: 70 to below 80
- D: 60 to below 70
- F: Below 60

A passing grade requires a C or better.

## Running Tests

From the repository root:

```bash
cd grade-calculator
go test ./...
```

To see individual test results:

```bash
go test -v ./...
```

## Code Coverage

Run the tests and generate a coverage report:

```bash
go test -coverprofile=coverage.out ./...
```

View the coverage percentage for each function:

```bash
go tool cover -func=coverage.out
```

## Main Functions

- `NewGradeCalculator()`: Creates a grade calculator.
- `AddGrade()`: Adds an assignment, exam, or essay grade.
- `GetFinalGrade()`: Returns the final letter grade.
- `calculateNumericalGrade()`: Calculates the weighted numerical grade.
- `computeAverage()`: Calculates the average of a list of grades.

## Notes

The original project includes tests that need to be reviewed against the requirements. Run the provided tests before making code changes.
# ESEP Grade Calculator

This library is a simple utility to calculate your grades.

Runtime requirements:
Golang 1.24+

Run unit tests:
```
cd grade_calculator/
go test ./...
```

Run unit tests with code coverage:
```
cd grade_calculator/
go test -coverprofile=coverage.txt ./...
```
