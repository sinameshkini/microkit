package models

import "testing"

func TestAmount_PersianString(t *testing.T) {
	tests := []struct {
		name string
		a    Amount
		want string
	}{
		{
			name: "TC01_normal_input",
			a:    100000,
			want: "۱۰۰٬۰۰۰",
		},
		{
			name: "TC02_zero_input",
			a:    0,
			want: "۰",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.PersianString(); got != tt.want {
				t.Errorf("PersianString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAmount_String(t *testing.T) {
	tests := []struct {
		name string
		a    Amount
		want string
	}{
		{
			name: "TC01_normal_input",
			a:    100000,
			want: "100000",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseAmount(t *testing.T) {
	type args struct {
		in string
	}
	tests := []struct {
		name string
		args args
		want Amount
	}{
		{
			name: "TC01_normal_input",
			args: args{
				in: "100000",
			},
			want: 100000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseAmount(tt.args.in); got != tt.want {
				t.Errorf("ParseAmount() = %v, want %v", got, tt.want)
			}
		})
	}
}
