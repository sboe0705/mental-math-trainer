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

	mistakes := 0
	startTime := time.Now().UnixMilli()
	for range *count {
		handleTask(*limit, operations, &mistakes)
	}
	endTime := time.Now().UnixMilli()

	fmt.Printf("You solved %d tasks in %d seconds!\n", *count, (endTime-startTime)/1000)
	fmt.Printf("You have made %d mistakes.\n", mistakes)
}

func handleTask(limit int, operations []string, mistakes *int) {
	task := NewTask(limit, operations)
	for {
		fmt.Print(task.Challenge())
		answer := readInteger()
		if answer == task.Result() {
			return
		}
		fmt.Println("Wrong!!! Repeat ...")
		*mistakes++
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
