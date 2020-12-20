package main

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_howSumTabulation(t *testing.T) {
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
			want: []int{2, 2, 3},
		},
		{
			args: args{
				targetSum: 7,
				numbers:   []int{2, 4},
			},
			want: nil,
		},
		{
			args: args{
				targetSum: 8,
				numbers:   []int{2, 3, 5},
			},
			want: []int{3, 5},
		},
		{
			args: args{
				targetSum: 300,
				numbers:   []int{7, 14},
			},
			want: nil,
		},
		{
			args: args{
				targetSum: 100,
				numbers:   []int{1, 2, 5, 25},
			},
			want: []int{25, 25, 25, 25},
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("case-%d", i), func(t *testing.T) {
			if got := howSumTabulation(tt.args.targetSum, tt.args.numbers); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("howSumTabulation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_bestSumTabulation(t *testing.T) {
	type args struct {
		targetSum int
		numbers   []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			args: args{
				targetSum: 7,
				numbers:   []int{5, 3, 4, 7},
			},
			want: []int{7},
		},
		{
			args: args{
				targetSum: 8,
				numbers:   []int{2, 3, 5},
			},
			want: []int{3, 5},
		},
		{
			args: args{
				targetSum: 8,
				numbers:   []int{1, 4, 5},
			},
			want: []int{4, 4},
		},
		{
			args: args{
				targetSum: 100,
				numbers:   []int{1, 2, 5, 25},
			},
			want: []int{25, 25, 25, 25},
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("case-%d", i), func(t *testing.T) {
			if got := bestSumTabulation(tt.args.targetSum, tt.args.numbers); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("bestSumTabulation() = %v, want %v", got, tt.want)
			}
		})
	}
}
