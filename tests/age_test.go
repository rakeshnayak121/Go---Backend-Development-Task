package tests

import (
	"testing"
	"user-api/utils"
)

func TestCalculateAge(t *testing.T) {

	age := utils.CalculateAge("2000-01-01")

	if age <= 0 {
		t.Errorf("Expected age > 0, got %d", age)
	}
}
