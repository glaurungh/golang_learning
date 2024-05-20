package questions

import (
	"encoding/json"
	"fmt"
	"os"
)

type Questions struct {
	Country string `json:"country"`
	Capital string `json:"capital"`
}

func LoadQuestions() ([]Questions, error) {
	jsonData, err := os.ReadFile("quiz.json")
	if err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}
	var questions []Questions
	err = json.Unmarshal(jsonData, &questions)
	if err != nil {
		return nil, fmt.Errorf("error parsing json data: %w", err)
	}
	return questions, nil
}
