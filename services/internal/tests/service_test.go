package tests

import (
	"testing"

	"github.com/felipemacedo1/go-antifraud-ms/internal/app"
)

func TestAdd(t *testing.T) {
	service := app.NewService()

	result := service.Add(2, 3)
	expected := 5

	if result != expected {
		t.Errorf("Add(2, 3) = %d; want %d", result, expected)
	}
}
