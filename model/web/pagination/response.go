package pagination

import (
	"encoding/json"
	"fmt"
)

type Response struct {
	Data      interface{} `json:"data"`
	TotalData int         `json:"total_data"`
}

func ResponseToJson(response Response) (map[string]interface{}, error) {
	jsonBytes, err := json.Marshal(response)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal response to JSON: %w", err)
	}

	var result map[string]interface{}
	err = json.Unmarshal(jsonBytes, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON bytes to map: %w", err)
	}

	return result, nil
}
