package employee

import (
	"errors"
)

// Employee struct
type Employee struct {
	name              string
	surname           string
	yearsOfEmployment int
}

// New function creates and returns a new instance of employee
func New(name string, surname string, yearsOfEmployment int) (*Employee, error) {
	if name == "" {
		return nil, errors.New("name must not be empty")
	}
	if surname == "" {
		return nil, errors.New("surname must not be empty")
	}
	if yearsOfEmployment < 0 {
		return nil, errors.New("yearsOfEmployment must not be non-negative number")
	}
	return &Employee{name, surname, yearsOfEmployment}, nil
}

func (e *Employee) calculateSalary(coef float32) float32 {
	return (float32)(e.yearsOfEmployment) * coef
}

// Name function
func (e *Employee) Name() string {
	return e.name
}

// Surname function
func (e *Employee) Surname() string {
	return e.surname
}

// PrintSalary function
func (e *Employee) PrintSalary() {
	println("Salary for:", e.name, e.surname, ":", e.calculateSalary(1.1))
}
