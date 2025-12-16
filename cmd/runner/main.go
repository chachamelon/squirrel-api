package main

import (
	"fmt"
	"os"

	"api-test-engine/internal/executor"
	"api-test-engine/internal/parser"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: runner <testcase.yaml or directory>")
		return
	}

	testCases, err := parser.LoadTestCases(os.Args[1])
	if err != nil {
		panic(err)
	}

	for _, tc := range testCases {
		result := executor.Run(tc)

		fmt.Printf("Test: %s\nPassed: %v\nDuration: %dms\n",
			result.Name, result.Passed, result.Duration)

		if !result.Passed {
			fmt.Println("Errors:")
			for _, e := range result.Errors {
				fmt.Println("-", e)
			}
		}
		fmt.Println("---")
	}
}
