package bestsum

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_bestsum(t *testing.T) {
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
				numbers:   []int{5, 3, 4, 7},
			},
			want: []int{7},
		},
		{
			args: args{
				targetSum: 8,
				numbers:   []int{2, 3, 5},
			},
			want: []int{5, 3},
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
			m := make(memo)
			if got := m.recursive(tt.args.targetSum, tt.args.numbers); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("memo.recursive() = %v, want %v", got, tt.want)
			}

			if got := tabulation(tt.args.targetSum, tt.args.numbers); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("tabulation() = %v, want %v", got, tt.want)
			}
		})
	}
}
