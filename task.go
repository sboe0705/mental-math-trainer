package main

import (
	"fmt"
	"math"
	"math/rand"
)

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

func NewTask(limit int) Task {
	operation := randomOperation()
	return NewTaskWithOperation(operation, limit)
}

func NewTaskWithOperation(operation string, limit int) Task {
	switch operation {
	case "+":
		return newAdditionTask(limit)
	case "-":
		return newSubtractionTask(limit)
	case "*", "x":
		return newMultiplicationTask(limit)
	case "/", ":":
		return newDivisionTask(limit)
	default:
	}
	panic(fmt.Sprintf("unknown operation '%s'", operation))
}

func randomOperation() string {
	ops := []string{"+", "-", "x", "/"}
	return ops[rand.Intn(len(ops))]
}

func newAdditionTask(limit int) *taskImpl {
	result := rand.Intn(limit + 1)
	number1 := rand.Intn(result + 1)
	number2 := result - number1
	return &taskImpl{number1: number1, operation: "+", number2: number2, result: result}
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
