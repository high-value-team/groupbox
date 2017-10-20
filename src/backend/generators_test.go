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
