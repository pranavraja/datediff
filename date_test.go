package main

import "testing"

func TestDate_Diff(t *testing.T) {
	tests := []struct {
		from         string
		to           string
		expectedDiff int
	}{
		{"2/6/1983", "22/6/1983", 19},
		{"4/7/1984", "25/12/1984", 173},
		{"1/3/1989", "3/8/1983", 2036},
	}
	for _, test := range tests {
		t.Run(test.from+" to "+test.to, func(t *testing.T) {
			from, _ := parseDate(test.from)
			to, _ := parseDate(test.to)
			actual := from.Diff(to)
			if actual != test.expectedDiff {
				t.Errorf("got %d days, wanted %d days", actual, test.expectedDiff)
			}
		})
	}
}
