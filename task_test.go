package main

import (
	"slices"
	"testing"
)

const limit = 100
const invocations = 1000

func TestCheckOperations(t *testing.T) {
	if CheckOperations(ValidOperations()) != nil {
		t.Errorf("correct operation is considered invalid")
	}
	if CheckOperations([]string{":"}) == nil {
		t.Errorf("invalid operation has not been detected")
	}
}

func TestNewTask(t *testing.T) {
	operations := []string{"+", "-", "x", "/"}
	for range invocations {
		task := NewTask(limit, operations)
		assertLimits(t, limit, task.(*taskImpl))
		assertValidOperation(t, task.(*taskImpl).operation, operations)
	}
}

func TestRandomOperation(t *testing.T) {
	operations := []string{"+", "-", "x", "/"}
	for range invocations {
		operation := randomOperation(operations)
		assertValidOperation(t, operation, operations)
	}
}

func assertValidOperation(t *testing.T, operation string, validOperations []string) {
	if !slices.Contains(validOperations, operation) {
		t.Fatalf("invalid operation '%s'", operation)
	}
}

func TestRandomNumber(t *testing.T) {
	min := 5
	max := 10
	for range invocations {
		number := randomNumber(min, max)
		if number < min || number > max {
			t.Fatalf("number %d is not between %d and %d (inclusive)", number, min, max)
		}
	}
}

func TestNewAdditionTask(t *testing.T) {
	for range invocations {
		task := newAdditionTask(limit)
		assertOperation(t, "+", task)
		assertLimits(t, limit, task)
		if task.number1+task.number2 != task.result {
			t.Fatalf("result %d is not the sum of %d and %d", task.result, task.number1, task.number2)
		}
	}
}

func TestNewSubtractionTask(t *testing.T) {
	for range invocations {
		task := newSubtractionTask(limit)
		assertOperation(t, "-", task)
		assertLimits(t, limit, task)
		if task.number1-task.number2 != task.result {
			t.Fatalf("result %d is not the difference of %d and %d", task.result, task.number1, task.number2)
		}
	}
}

func TestNewMultiplicationTask(t *testing.T) {
	for range invocations {
		task := newMultiplicationTask(limit)
		assertOperation(t, "x", task)
		assertLimits(t, limit, task)
		if task.number1*task.number2 != task.result {
			t.Fatalf("result %d is not the product of %d and %d", task.result, task.number1, task.number2)
		}
	}
}

func TestNewDivisionTask(t *testing.T) {
	for range invocations {
		task := newDivisionTask(limit)
		assertOperation(t, "/", task)
		assertLimits(t, limit, task)
		if task.number1/task.number2 != task.result {
			t.Fatalf("result %d is not the division of %d by %d", task.result, task.number1, task.number2)
		}
	}
}

func assertOperation(t *testing.T, expectedOperation string, task *taskImpl) {
	if task.operation != expectedOperation {
		t.Fatalf("wrong operation (expected '%s', but was '%s')", expectedOperation, task.operation)
	}
}

func assertLimits(t *testing.T, limit int, task *taskImpl) {
	if task.number1 > limit || task.number2 > limit || task.result > limit {
		t.Fatalf("limit of %d exceeded in task '%s%d'", limit, task.Challenge(), task.result)
	}
	if task.number1 < 2 || task.number2 < 2 || task.result < 2 {
		t.Fatalf("lower limit of 2 was not reached in task '%s%d'", task.Challenge(), task.result)
	}
}
