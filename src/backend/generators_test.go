package backend

import (
	"testing"
	"fmt"
)

func TestGenerateKey(t *testing.T) {
	result1 := generateKey()
	result2 := generateKey()

	fmt.Printf("Key generated: %v\n", result1)

	if result1 == result2 {
		t.Error("Key generation failed! Two matching keys were generated.")
	}
}


func TestNewNicknameGenerator(t *testing.T) {
	sut := NewNicknameGenerator_explicit([]string{"a", "b"}, []string{"x", "y"})

	var results []string
	results = append(results, sut.next())
	results = append(results, sut.next())
	results = append(results, sut.next())
	results = append(results, sut.next())
	results = append(results, sut.next())
	results = append(results, sut.next())

	uniqueResults := map[string]string {}
	for _,r := range results {
		fmt.Printf("Nickname generated: %s\n", r)

		if _,ok := uniqueResults[r]; ok {
			t.Error("Duplicate key nickname generated!")
			return
		}
		uniqueResults[r] = ""
	}
}
