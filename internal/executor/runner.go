package executor

import (
	"time"

	"fmt"

	"github.com/chachamelon/squirrel-api/internal/assert"
	"github.com/chachamelon/squirrel-api/internal/model"
)

func Run(tc *model.TestCase) *model.TestResult {
	start := time.Now()
	result := &model.TestResult{
		Name:   tc.Name,
		Passed: true,
	}

	resp, body, err := ExecuteRequest(
		tc.Request.Method,
		tc.Request.URL,
		tc.Request.Headers,
		tc.Request.Body,
	)
	fmt.Println(string(body))

	if err != nil {
		result.Passed = false
		result.Errors = append(result.Errors, err.Error())
		return result
	}

	if err := assert.Status(resp.StatusCode, tc.Assert.Status); err != nil {
		result.Passed = false
		result.Errors = append(result.Errors, err.Error())
	}

	for field, rule := range tc.Assert.Json {
		if rule == "not_null" {
			if err := assert.JsonNotNull(body, field); err != nil {
				result.Passed = false
				result.Errors = append(result.Errors, err.Error())
			}
		}
	}

	result.Duration = time.Since(start).Milliseconds()
	return result
}
