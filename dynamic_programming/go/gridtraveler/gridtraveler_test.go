package gridtraveler

import (
	"fmt"
	"testing"
)

func Test_gridtraveler(t *testing.T) {
	type args struct {
		m int
		n int
	}
	tests := []struct {
		args args
		want int
	}{
		{
			args: args{
				m: 3,
				n: 3,
			},
			want: 6,
		},
		{
			args: args{
				m: 2,
				n: 3,
			},
			want: 3,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("case-%d", i), func(t *testing.T) {
			if got := recursive(tt.args.m, tt.args.n); got != tt.want {
				t.Errorf("recursive() = %v, want %v", got, tt.want)
			}

			m := make(memo)
			if got := m.recursive(tt.args.m, tt.args.n); got != tt.want {
				t.Errorf("memo.recursive() = %v, want %v", got, tt.want)
			}

			if got := tabulation(tt.args.m, tt.args.n); got != tt.want {
				t.Errorf("tabulation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_gridtraveler_cost(t *testing.T) {
	type args struct {
		m int
		n int
	}
	tests := []struct {
		args args
		want int
	}{
		{
			args: args{
				m: 18,
				n: 18,
			},
			want: 2333606220,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("case-%d", i), func(t *testing.T) {
			m := make(memo)
			if got := m.recursive(tt.args.m, tt.args.n); got != tt.want {
				t.Errorf("memo.recursive() = %v, want %v", got, tt.want)
			}

			if got := tabulation(tt.args.m, tt.args.n); got != tt.want {
				t.Errorf("tabulation() = %v, want %v", got, tt.want)
			}
		})
	}
}
