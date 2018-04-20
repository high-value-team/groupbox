// +build unit

package interior_interactions

import (
	"fmt"
	"testing"
)

func TestGenerateKey(t *testing.T) {
	result1 := GenerateKey()
	result2 := GenerateKey()

	fmt.Printf("Key generated: %v\n", result1)

	if result1 == result2 {
		t.Error("Key generation failed! Two matching keys were generated.")
	}
}

func TestNewNicknameGenerator(t *testing.T) {
	sut := NewNicknameGenerator()

	var results []string
	results = append(results, sut.Next())
	results = append(results, sut.Next())
	results = append(results, sut.Next())
	results = append(results, sut.Next())
	results = append(results, sut.Next())
	results = append(results, sut.Next())

	uniqueResults := map[string]string{}
	for _, r := range results {
		fmt.Printf("Nickname generated: %s\n", r)

		if _, ok := uniqueResults[r]; ok {
			t.Error("Duplicate nickname generated!")
			return
		}
		uniqueResults[r] = ""
	}
}
