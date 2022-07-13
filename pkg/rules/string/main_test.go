package rule_string

import (
	"testing"
)

func TestEquals(t *testing.T) {
	type args struct {
		s      string
		length int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{"Правильный email", args{s: "qwe123", length: 6}, true},
		{"Ошибочный email", args{s: "qwe123", length: 5}, false},
		{"Правильный email", args{s: "ыфв7ук", length: 6}, true},
		{"Ошибочный email", args{s: "ыфв7ук", length: 5}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Equals(tt.args.s, tt.args.length); got != tt.want {
				t.Errorf("Equals() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetSymbolPosition(t *testing.T) {
	type args struct {
		s      string
		symbol rune
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"Правильный", args{s: "q@we123", symbol: '@'}, 1},
		{"Правильный", args{s: "qwe123", symbol: '5'}, -1},
		{"Правильный", args{s: "ыфв7ук", symbol: 'в'}, 2},
		{"Правильный", args{s: "ыфв7ук", symbol: 'ы'}, 0},
		{"Правильный", args{s: "фв7ыук", symbol: 'ы'}, 3},
		{"Правильный", args{s: "ыфв7ук", symbol: '№'}, -1},
		{"Правильный", args{s: "ыфв7ук", symbol: 'у'}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetSymbolPosition(tt.args.s, tt.args.symbol); got != tt.want {
				t.Errorf("GetSymbolPosition() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetLength(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"Правильный", args{s: "q@we123"}, 7},
		{"Правильный", args{s: "фывф"}, 4},
		{"Правильный", args{s: ""}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetLength(tt.args.s); got != tt.want {
				t.Errorf("GetLength() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSubstr(t *testing.T) {
	type args struct {
		input string
		start int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"Правильный", args{input: "q@we123", start: 1}, "@we123"},
		{"Правильный", args{input: "фывф", start: 3}, "ф"},
		{"Правильный", args{input: "фывф", start: 6}, ""},
		{"Правильный", args{input: "", start: 0}, ""},
		{"Правильный", args{input: "", start: 1}, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Substr(tt.args.input, tt.args.start); got != tt.want {
				t.Errorf("Substr() = %v, want %v", got, tt.want)
			}
		})
	}
}
