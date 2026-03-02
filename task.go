package main

import (
	"fmt"
	"math"
	"math/rand"
	"slices"
)

func ValidOperations() []string {
	return []string{"+", "-", "x", "/"}
}

func CheckOperations(operations []string) error {
	validOperations := ValidOperations()
	for _, operation := range operations {
		if !slices.Contains(validOperations, operation) {
			return fmt.Errorf("invalid operation '%s'", operation)
		}
	}
	return nil
}

type Task interface {
	Challenge() string
	Result() int
}

type taskImpl struct {
	operation string
	number1   int
	number2   int
	result    int
}

func (this *taskImpl) Challenge() string {
	return fmt.Sprintf("%d %s %d = ", this.number1, this.operation, this.number2)
}

func (this *taskImpl) Result() int {
	return this.result
}

func NewTask(limit int, operations []string) Task {
	operation := randomOperation(operations)
	return NewTaskWithOperation(operation, limit)
}

func NewTaskWithOperation(operation string, limit int) Task {
	switch operation {
	case "+":
		return newAdditionTask(limit)
	case "-":
		return newSubtractionTask(limit)
	case "x":
		return newMultiplicationTask(limit)
	case "/":
		return newDivisionTask(limit)
	default:
	}
	panic(fmt.Sprintf("unknown operation '%s'", operation))
}

func randomOperation(operations []string) string {
	return operations[rand.Intn(len(operations))]
}

func newAdditionTask(limit int) *taskImpl {
	result := randomNumber(4, limit)
	number1 := randomNumber(2, result-2)
	number2 := result - number1
	return &taskImpl{number1: number1, operation: "+", number2: number2, result: result}
}

func randomNumber(min, max int) int {
	return min + rand.Intn(max+1-min)
}

func newSubtractionTask(limit int) *taskImpl {
	additionTask := newAdditionTask(limit)
	return &taskImpl{number1: additionTask.result, operation: "-", number2: additionTask.number1, result: additionTask.number2}
}

func newMultiplicationTask(limit int) *taskImpl {
	sqrt := int(math.Sqrt(float64(limit)))
	number1 := 2 + rand.Intn(sqrt-1)
	number2 := 2 + rand.Intn(sqrt-1)
	result := number1 * number2
	return &taskImpl{number1: number1, operation: "x", number2: number2, result: result}
}

func newDivisionTask(limit int) *taskImpl {
	multiplicationTask := newMultiplicationTask(limit)
	return &taskImpl{number1: multiplicationTask.result, operation: "/", number2: multiplicationTask.number1, result: multiplicationTask.number2}
}
