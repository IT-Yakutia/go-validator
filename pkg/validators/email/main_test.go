package validate_email

import (
	"testing"
)

func TestEmail_Validate(t *testing.T) {
	type fields struct {
		email string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
		{"Правильный email", fields{email: "a@d.s"}, true},
		{"Правильный email", fields{email: "a.a@d.s.s"}, true},
		{"Правильный email", fields{email: "a-a@d-d.s"}, true},
		{"Правильный email", fields{email: "фывфы@фыв.фыв"}, true},
		{"Ошибочный email", fields{email: "a@d@s"}, false},
		{"Ошибочный email", fields{email: "a.d.s"}, false},
		{"Ошибочный email", fields{email: "a.d@s"}, false},
		{"Ошибочный email", fields{email: "@d.s"}, false},
		{"Ошибочный email", fields{email: "a@.s"}, false},
		{"Ошибочный email", fields{email: "a@d."}, false},
		{"Ошибочный email", fields{email: "a.d"}, false},
		{"Ошибочный email", fields{email: "a@d"}, false},
		{"Ошибочный email", fields{email: ".@"}, false},
		{"Ошибочный email", fields{email: "a@"}, false},
		{"Ошибочный email", fields{email: "@d"}, false},
		{"Ошибочный email", fields{email: "d."}, false},
		{"Ошибочный email", fields{email: ".s"}, false},
		{"Ошибочный email", fields{email: "@."}, false},
		{"Ошибочный email", fields{email: "a"}, false},
		{"Ошибочный email", fields{email: "."}, false},
		{"Ошибочный email", fields{email: "@"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Email{
				email: tt.fields.email,
			}
			if got := e.Validate(); got != tt.want {
				t.Errorf("Email.Validate() = %v, want %v", got, tt.want)
			}
		})
	}
}
