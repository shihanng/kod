package main

import (
	"fmt"
	"reflect"
	"testing"
)

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
