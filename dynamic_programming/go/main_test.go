package main

import (
	"fmt"
	"reflect"
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

func Test_howSum(t *testing.T) {
	type args struct {
		targetSum int
		numbers   []int
	}
	tests := []struct {
		args args
		want []int
	}{
		{
			args: args{
				targetSum: 7,
				numbers:   []int{2, 3},
			},
			want: []int{3, 2, 2},
		},
		{
			args: args{
				targetSum: 7,
				numbers:   []int{2, 4},
			},
			want: nil,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("case-%d", i), func(t *testing.T) {
			if got := howSum(tt.args.targetSum, tt.args.numbers); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("howSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fastHowSum_howSum(t *testing.T) {
	type args struct {
		targetSum int
		numbers   []int
	}
	tests := []struct {
		s    fastHowSum
		args args
		want []int
	}{
		{
			s: make(fastHowSum),
			args: args{
				targetSum: 7,
				numbers:   []int{2, 4},
			},
			want: nil,
		},
		{
			s: make(fastHowSum),
			args: args{
				targetSum: 8,
				numbers:   []int{2, 3, 5},
			},
			want: []int{2, 2, 2, 2},
		},
		{
			s: make(fastHowSum),
			args: args{
				targetSum: 300,
				numbers:   []int{7, 14},
			},
			want: nil,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("case-%d", i), func(t *testing.T) {
			if got := tt.s.howSum(tt.args.targetSum, tt.args.numbers); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fastHowSum.howSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_bestSum_bestSum(t *testing.T) {
	type args struct {
		targetSum int
		numbers   []int
	}
	tests := []struct {
		s    bestSum
		args args
		want []int
	}{
		{
			s: make(bestSum),
			args: args{
				targetSum: 7,
				numbers:   []int{5, 3, 4, 7},
			},
			want: []int{7},
		},
		{
			s: make(bestSum),
			args: args{
				targetSum: 8,
				numbers:   []int{2, 3, 5},
			},
			want: []int{5, 3},
		},
		{
			s: make(bestSum),
			args: args{
				targetSum: 8,
				numbers:   []int{1, 4, 5},
			},
			want: []int{4, 4},
		},
		{
			s: make(bestSum),
			args: args{
				targetSum: 100,
				numbers:   []int{1, 2, 5, 25},
			},
			want: []int{25, 25, 25, 25},
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("case-%d", i), func(t *testing.T) {
			if got := tt.s.bestSum(tt.args.targetSum, tt.args.numbers); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("bestSum.bestSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
