package cansum

import (
	"fmt"
	"testing"
)

func Test_cansum(t *testing.T) {
	type args struct {
		targetSum int
		numbers   []int
	}
	tests := []struct {
		args args
		want bool
	}{
		{
			args: args{
				targetSum: 7,
				numbers:   []int{2, 3},
			},
			want: true,
		},
		{
			args: args{
				targetSum: 7,
				numbers:   []int{5, 3, 4, 7},
			},
			want: true,
		},
		{
			args: args{
				targetSum: 7,
				numbers:   []int{2, 4},
			},
			want: false,
		},
		{
			args: args{
				targetSum: 8,
				numbers:   []int{2, 3, 5},
			},
			want: true,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("case-%d", i), func(t *testing.T) {
			if got := recursive(tt.args.targetSum, tt.args.numbers); got != tt.want {
				t.Errorf("recursive() = %v, want %v", got, tt.want)
			}

			m := make(memo)
			if got := m.recursive(tt.args.targetSum, tt.args.numbers); got != tt.want {
				t.Errorf("memo.recursive() = %v, want %v", got, tt.want)
			}

			if got := tabulation(tt.args.targetSum, tt.args.numbers); got != tt.want {
				t.Errorf("tabulation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_cansum_cost(t *testing.T) {
	type args struct {
		targetSum int
		numbers   []int
	}
	tests := []struct {
		args args
		want bool
	}{
		{
			args: args{
				targetSum: 300,
				numbers:   []int{7, 14},
			},
			want: false,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("case-%d", i), func(t *testing.T) {
			m := make(memo)
			if got := m.recursive(tt.args.targetSum, tt.args.numbers); got != tt.want {
				t.Errorf("memo.recursive() = %v, want %v", got, tt.want)
			}

			if got := tabulation(tt.args.targetSum, tt.args.numbers); got != tt.want {
				t.Errorf("tabulation() = %v, want %v", got, tt.want)
			}
		})
	}
}
