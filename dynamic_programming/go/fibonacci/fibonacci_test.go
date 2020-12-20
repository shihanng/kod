package fibonacci

import (
	"fmt"
	"testing"
)

func Test_fibonacci(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		args args
		want int
	}{
		{
			args: args{
				n: 6,
			},
			want: 8,
		},
		{
			args: args{
				n: 7,
			},
			want: 13,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("case-%d", i), func(t *testing.T) {
			if got := recursive(tt.args.n); got != tt.want {
				t.Errorf("recursive() = %v, want %v", got, tt.want)
			}

			m := make(memo)
			if got := m.recursive(tt.args.n); got != tt.want {
				t.Errorf("memo.recursive() = %v, want %v", got, tt.want)
			}

			if got := tabulation(tt.args.n); got != tt.want {
				t.Errorf("tabulation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fibonacci_cost(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		args args
		want int
	}{
		{
			args: args{
				n: 50,
			},
			want: 12586269025,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("case-%d", i), func(t *testing.T) {
			m := make(memo)
			if got := m.recursive(tt.args.n); got != tt.want {
				t.Errorf("memo.recursive() = %v, want %v", got, tt.want)
			}

			if got := tabulation(tt.args.n); got != tt.want {
				t.Errorf("tabulation() = %v, want %v", got, tt.want)
			}
		})
	}
}
