package handlers

import (
	"bytes"
	"strings"
	"testing"
)

func TestInputHandler_ValidInput(t *testing.T) {
	input := "1 2 3 4 5\n"
	expected := []int{1, 2, 3, 4, 5}

	mockInput := bytes.NewBufferString(input)
	result, err := InputHandler(3, mockInput)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !equalSlices(result, expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

func TestInputHandler_InvalidInput(t *testing.T) {
	input := "1 2 three 4\n"
	mockInput := bytes.NewBufferString(input)

	_, err := InputHandler(3, mockInput)
	if err == nil || !strings.Contains(err.Error(), "please enter integers only") {
		t.Errorf("expected integer error, got %v", err)
	}
}

func TestInputHandler_NotEnoughNumbers(t *testing.T) {
	input := "1 2\n"
	mockInput := bytes.NewBufferString(input)

	_, err := InputHandler(3, mockInput)
	if err == nil || !strings.Contains(err.Error(), "please enter minimum") {
		t.Errorf("expected minimum number error, got %v", err)
	}
}

func equalSlices(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
