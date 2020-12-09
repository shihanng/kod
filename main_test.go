package main

import (
	"fmt"
	"testing"
)

func Test_canSum(t *testing.T) {
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
			if got := canSum(tt.args.targetSum, tt.args.numbers); got != tt.want {
				t.Errorf("canSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fastCanSum_canSum(t *testing.T) {
	type args struct {
		targetSum int
		numbers   []int
	}
	tests := []struct {
		c    fastCanSum
		args args
		want bool
	}{
		{
			c: make(fastCanSum),
			args: args{
				targetSum: 7,
				numbers:   []int{2, 3},
			},
			want: true,
		},
		{
			c: make(fastCanSum),
			args: args{
				targetSum: 300,
				numbers:   []int{7, 14},
			},
			want: false,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("case-%d", i), func(t *testing.T) {
			if got := tt.c.canSum(tt.args.targetSum, tt.args.numbers); got != tt.want {
				t.Errorf("fastCanSum.canSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
