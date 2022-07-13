package rule_number

import "testing"

func TestIsNumberWithLenght(t *testing.T) {
	type args struct {
		n      string
		length int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{"Правильный", args{n: "фывф", length: 3}, false},
		{"Правильный", args{n: "ф32ф", length: 4}, false},
		{"Правильный", args{n: "234234", length: 6}, true},
		{"Правильный", args{n: "23", length: 2}, true},
		{"Правильный", args{n: "", length: 0}, false},
		{"Правильный", args{n: "234234", length: 5}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsNumberWithLenght(tt.args.n, tt.args.length); got != tt.want {
				t.Errorf("IsNumberWithLenght() = %v, want %v", got, tt.want)
			}
		})
	}
}
