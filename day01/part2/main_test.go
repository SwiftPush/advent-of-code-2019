package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFoo(t *testing.T) {
	tests := []struct {
		mass         int
		expectedFuel int
	}{
		{
			119965,
			39986,
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("test_%d", i), func(t *testing.T) {
			actualFuel := calculateFuelRequired(test.mass)
			assert.Equal(t, test.expectedFuel, actualFuel)
		})
	}
}
