package main

import (
	"testing"
)

func Test_GenerateChecksum(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  int
	}{
		{
			name:  "want 0",
			input: []string{"abcdef"},
			want:  0,
		},
		{
			name:  "want 1",
			input: []string{"bababc"},
			want:  1,
		},
		{
			name:  "want 2",
			input: []string{"bababc", "abbcdde"},
			want:  2,
		},
		{
			name: "want 3",
			input: []string{
				"oiwcdpbseqgxryfmlpktnupvza", // 0 | 1
				"oiwddpbsuqhxryfmlgkznujvza", // 1 | 0
				"ziwcdpbsechxrvfmlgktnujvza", // 1 | 0
				"oiwcgpbseqhxryfmmgktnhjvza", // 1 | 0
			},
			want: 3, // 3 * 1
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GenerateChecksum(tt.input)

			if got != tt.want {
				t.Errorf("got: %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_Basic(t *testing.T) {
	input := []string{
		"oiwcdpbseqgxryfmlpktnupvza", // 0 | 1
		"oiwddpbsuqhxryfmlgkznujvza", // 1 | 0
		"ziwcdpbsechxrvfmlgktnujvza", // 1 | 0
		"oiwcgpbseqhxryfmmgktnhjvza", // 1 | 0
	}
	got := GenerateChecksum(input)

	if got != 3 {
		t.Errorf("got: %v, want %v", got, 3)
	}
}

func BenchmarkGenerateChecksum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		input := []string{
			"oiwcdpbseqgxryfmlpktnupvza", // 0 | 1
			"oiwddpbsuqhxryfmlgkznujvza", // 1 | 0
			"ziwcdpbsechxrvfmlgktnujvza", // 1 | 0
			"oiwcgpbseqhxryfmmgktnhjvza", // 1 | 0
		}
		_ = GenerateChecksum(input)
	}
}

// func TestMain(m *testing.M) {
// 	result := testing.Benchmark(BenchmarkGenerateChecksum)

// 	fmt.Println("total iterations:", result.N)
// 	fmt.Println("----------")
// 	fmt.Println("allocs/op:", result.AllocsPerOp())
// 	fmt.Println("total allocs:", result.MemAllocs)
// 	fmt.Println("----------")
// 	fmt.Println("bytes/op:", result.AllocedBytesPerOp())
// 	fmt.Println("total bytes:", result.MemBytes)
// 	fmt.Println("----------")
// 	fmt.Println("bytes:", result.Bytes)
// 	fmt.Println("----------")
// 	fmt.Println("ns/op:", result.NsPerOp())
// }
