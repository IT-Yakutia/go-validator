package validate_phone

import (
	"testing"
)

func TestPhone_GetNumber(t *testing.T) {
	type fields struct {
		code   string
		number string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
		{"Правильный", fields{code: "+7", number: "234234234"}, "+7234234234"},
		{"Правильный", fields{code: "8", number: "4234"}, "84234"},
		{"Правильный", fields{code: "7", number: "123"}, "7123"},
		{"Правильный", fields{code: "sdf", number: "sdf"}, "sdfsdf"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Phone{
				code:   tt.fields.code,
				number: tt.fields.number,
			}
			if got := p.GetNumber(); got != tt.want {
				t.Errorf("Phone.GetNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPhone_ValidateRussian(t *testing.T) {
	type fields struct {
		code   string
		number string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
		{"Правильный", fields{code: "+7", number: "234234234"}, false},
		{"Правильный", fields{code: "+798123", number: "324234"}, false},
		{"Правильный", fields{code: "+7", number: "9142736836"}, true},
		{"Правильный", fields{code: "8", number: "9142736836"}, true},
		{"Правильный", fields{code: "7", number: "9142736836"}, true},
		{"Правильный", fields{code: "77", number: "142736836"}, false},
		{"Правильный", fields{code: "79", number: "142736836"}, false},
		{"Правильный", fields{code: "7", number: "91427368360"}, false},
		{"Правильный", fields{code: "7", number: "(914)273-68-36"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Phone{
				code:   tt.fields.code,
				number: tt.fields.number,
			}
			if got := p.ValidateRussian(); got != tt.want {
				t.Errorf("Phone.ValidateRussian() = %v, want %v", got, tt.want)
			}
		})
	}
}
