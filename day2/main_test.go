package main

import "testing"

func Test_GenerateChecksum(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  int
	}{
		{
			name:  "want 1",
			input: []string{"abcdef", "bababc", "abbcde"},
			want:  1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GenerateChecksum(tt.input)
			if err != nil {
				t.Errorf("unexpected error: %s", err.Error())
			}

			if got != tt.want {
				t.Errorf("got: %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_CalcTotal(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  int
	}{
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
			got, err := CalcTotal(tt.input)
			if err != nil {
				t.Errorf("unexpected error: %s", err.Error())
			}

			if got != tt.want {
				t.Errorf("got: %v, want %v", got, tt.want)
			}
		})
	}
}
