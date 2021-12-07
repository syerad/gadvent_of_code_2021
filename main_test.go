package main

import (
	"testing"
)

func Test_countIncreasedSingals(t *testing.T) {
	tests := []struct {
		name        string
		signalArray []int16
		want        int
	}{
		{"validPostiveNumbers", provideAntennaSignals(), 7},
		{"validNegativeNumbers", provideNegativeAntennaSignals(), 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countIncreasedSignals(tt.signalArray); got != tt.want {
				t.Errorf("countIncreasedSingals() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_countIncreasedSignalsTriplet(t *testing.T) {
	tests := []struct {
		name        string
		signalArray []int16
		want        int
	}{
		{"validPostiveNumbers", provideAntennaSignals(), 5},
		{"validNegativeNumbers", provideNegativeAntennaSignals(), 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countIncreasedSignalsTriplet(tt.signalArray); got != tt.want {
				t.Errorf("countIncreasedSignalsTriplet() with %v = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}

func provideNegativeAntennaSignals() []int16 {
	return []int16{
		-1,
		-3,
		3,
		3,
		-5,
		-2,
	}
}

func provideAntennaSignals() []int16 {
	return []int16{
		199,
		200,
		208,
		210,
		200,
		207,
		240,
		269,
		260,
		263,
	}
}
