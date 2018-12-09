package main

import (
	"io"
	"strings"
	"testing"
)

func Test_CalcTotal(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  int
	}{
		{
			"want zero",
			[]string{""},
			0,
		},
		{
			"want three",
			[]string{"+1", "+1", "+1"},
			3,
		},
		{
			"want zero",
			[]string{"+1", "+1", "-2"},
			0,
		},
		{
			"want minus six",
			[]string{"-1", "-2", "-3"},
			-6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CalcTotal(tt.input...)
			if err != nil {
				t.Errorf("unexpected error: %s", err.Error())
			}

			if got != tt.want {
				t.Errorf("got: %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_CalcTotalReader(t *testing.T) {
	tests := []struct {
		name  string
		input io.Reader
		want  int
	}{
		{
			"want three",
			strings.NewReader("+1\n+1\n+1"),
			3,
		},
		{
			"want zero",
			strings.NewReader("+1\n+1\n-2"),
			0,
		},
		{
			"want minus six",
			strings.NewReader("-1\n-2\n-3"),
			-6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CalcTotalWithReader(tt.input)
			if err != nil {
				t.Errorf("unexpected error: %s", err.Error())
			}

			if got != tt.want {
				t.Errorf("got: %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_DupeFreq(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  int
	}{
		{
			"want zero",
			[]string{""},
			0,
		},
		{
			"want two",
			[]string{"+1", "-2", "+3", "+1"},
			2,
		},
		{
			"want ten",
			[]string{"+3", "+3", "+4", "-2", "-4"},
			10,
		},
		{
			"want five",
			[]string{"-6", "+3", "+8", "+5", "-6"},
			5,
		},
		{
			"want fourteen",
			[]string{"+7", "+7", "-2", "-7", "-4"},
			14,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := DuplicateFreq(tt.input)
			if got != tt.want {
				t.Errorf("got: %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkCalcTotal(b *testing.B) {
	input := []string{"+11234567", "+1123456", "+198765432"}

	for n := 0; n < b.N; n++ {
		_, err := CalcTotal(input...)
		if err != nil {
			b.Errorf("unexpected error: %s", err.Error())
		}
	}
}

func BenchmarkCalcTotalWithReader(b *testing.B) {
	input := strings.NewReader("+11234567\n+1123456\n+198765432")

	for n := 0; n < b.N; n++ {
		_, err := CalcTotalWithReader(input)
		if err != nil {
			b.Errorf("unexpected error: %s", err.Error())
		}
	}
}
