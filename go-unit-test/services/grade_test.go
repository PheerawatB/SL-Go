package services_test

import (
	"fmt"
	"gotest/services"
	"testing"
	// "github.com/stretchr/testify/assert"
)

func TestCheckGrade(t *testing.T) {

	type testCase struct {
		name     string
		score    int
		expected string
	}

	cases := []testCase{
		{"A", 80, "A"},
		{"B", 70, "B"},
		{"C", 60, "C"},
		{"D", 50, "D"},
		{"F", 40, "F"},
		{"F", 0, "F"},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			grade := services.CheckGrade(tc.score)
			// assert.Equal(t, tc.expected, grade)
			if grade != tc.expected {
				t.Errorf("expected %v, got %v", tc.expected, grade)
			}
		})
	}
}

func BenchmarkCheckGrade(b *testing.B) {
	for i := 0; i < b.N; i++ {
		services.CheckGrade(80)
	}
}

func ExampleCheckGrade() {
	grade := services.CheckGrade(80)
	fmt.Println(grade)
	// Output: A
}
