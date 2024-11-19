// challenges/testing/begin/main_test.go
package main

import "testing"

// write a test for letterCounter.count
func TestLetterCount(t *testing.T) {
	var tests = map[string]struct {
		input string
		want  int
	}{
		"empty": {input: "", want: 0},
		"one":   {input: "a2 32 ^ / &", want: 1},
		"two":   {input: "546 %ab *", want: 2},
	}

	lc := letterCounter{}
	for _, tc := range tests {
		t.Run(tc.input, func(t *testing.T) {
			if got := lc.count(tc.input); got != tc.want {
				t.Errorf("got %v, want %v", got, tc.want)
			}

		})
	}

}

// write a test for numberCounter.count
func TestNumberCount(t *testing.T) {
	var tests = map[string]struct {
		input string
		want  int
	}{
		"empty": {input: "", want: 0},
		"one":   {input: "a2 32 ^ / &", want: 3},
		"two":   {input: "546 %ab *", want: 3},
	}

	lc := numberCounter{}
	for _, tc := range tests {
		t.Run(tc.input, func(t *testing.T) {
			if got := lc.count(tc.input); got != tc.want {
				t.Errorf("got %v, want %v", got, tc.want)
			}

		})
	}

}

// write a test for symbolCounter.count
func TestSymbolCount(t *testing.T) {
	var tests = map[string]struct {
		input string
		want  int
	}{
		"empty": {input: "", want: 0},
		"one":   {input: "a2 32 ^ / &", want: 7}, // spaces are also symbols
		"two":   {input: "546 %ab *", want: 4},
	}

	lc := symbolCounter{}
	for _, tc := range tests {
		t.Run(tc.input, func(t *testing.T) {
			if got := lc.count(tc.input); got != tc.want {
				t.Errorf("got %v, want %v", got, tc.want)
			}

		})
	}
}
