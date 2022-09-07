package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrderBySimilarity(t *testing.T) {
	people := []Person{
		{"John Doe", 60},
		{"John Wick", 30},
		{"John Travolta", 40},
		{"John Jow", 50},
	}

	tests := []struct {
		testCase string
		source   string
		expected []Person
	}{
		{
			testCase: "order by John Travolta",
			source:   "John Travolta",
			expected: []Person{{"John Travolta", 40}, {"John Doe", 60}, {"John Jow", 50}, {"John Wick", 30}},
		},
		{
			testCase: "order by John Doe",
			source:   "John Doe",
			expected: []Person{{"John Doe", 60}, {"John Jow", 50}, {"John Wick", 30}, {"John Travolta", 40}},
		},
		{
			testCase: "order by John Wick",
			source:   "John Wick",
			expected: []Person{{"John Wick", 30}, {"John Doe", 60}, {"John Jow", 50}, {"John Travolta", 40}},
		},
		{
			testCase: "order by John Jow",
			source:   "John Jow",
			expected: []Person{{"John Jow", 50}, {"John Doe", 60}, {"John Wick", 30}, {"John Travolta", 40}},
		},
	}

	for _, tc := range tests {
		t.Run(tc.testCase, func(t *testing.T) {
			result := orderPeopleBySimilarityName(tc.source, people)
			assert.Equal(t, tc.expected, result)
		})
	}
}
