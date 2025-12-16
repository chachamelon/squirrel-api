package assert

import (
	"encoding/json"
	"fmt"
)

func JsonNotNull(body []byte, field string) error {
	var data map[string]interface{}
	json.Unmarshal(body, &data)

	if _, ok := data[field]; !ok {
		return fmt.Errorf("json field '%s' not found", field)
	}
	return nil
}
