package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/pflag"
)

func main() {
	count := pflag.IntP("count", "c", 20, "Number of problems to generate and solve.")
	limit := pflag.IntP("limit", "l", 100, "Maximum number used in generated problems.")
	operationsString := pflag.StringP("operations", "o", "+-x/", "Available math operations.")
	pflag.Parse()

	operations := make([]string, 0, len(*operationsString))
	for _, r := range *operationsString {
		operations = append(operations, string(r))
	}
	error := CheckOperations(operations)
	if error != nil {
		fmt.Println(error.Error())
		return
	}

	errors := 0
	startTime := time.Now().UnixMilli()
	for range *count {
		handleTask(*limit, operations, &errors)
	}
	endTime := time.Now().UnixMilli()
	
	totalTime := (endTime-startTime)/1000
	avgTime := float64(totalTime)/float64(*count)
	errorRate := float64(errors)/float64(*count)
	effectiveTime := avgTime*(1+errorRate*2)
	score := 1000/effectiveTime

	fmt.Printf("You solved %d tasks in %d seconds (%.1f seconds per task)!\n", *count, totalTime, avgTime)
	fmt.Printf("You have made %d errors.\n", errors)
	fmt.Printf("Your score: %d\n", score)
}

func handleTask(limit int, operations []string, errors *int) {
	task := NewTask(limit, operations)
	for {
		fmt.Print(task.Challenge())
		answer := readInteger()
		if answer == task.Result() {
			return
		}
		fmt.Println("Wrong!!! Repeat ...")
		*errors++
	}
}

func readInteger() int {
	reader := bufio.NewReader(os.Stdin)
	for {
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		value, err := strconv.Atoi(input)
		if err == nil {
			return value
		}
	}
}
