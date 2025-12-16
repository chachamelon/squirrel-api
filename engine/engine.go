package engine

import (
	"fmt"

	"github.com/chachamelon/squirrel-api/internal/executor"
	"github.com/chachamelon/squirrel-api/internal/parser"
)

func RunTest(path string) (string, error) {
	testCases, err := parser.LoadTestCases(path)
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

	return "hi", nil
}
