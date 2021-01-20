package combk

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_combk(t *testing.T) {
	type args struct {
		n int
		k int
	}
	tests := []struct {
		args args
		want [][]int
	}{
		{
			args: args{
				n: 1,
				k: 1,
			},
			want: [][]int{{1}},
		},
		{
			args: args{
				n: 4,
				k: 1,
			},
			want: [][]int{{4}, {1}, {2}, {3}},
		},
		{
			args: args{
				n: 4,
				k: 2,
			},
			want: [][]int{{1, 4}, {2, 4}, {1, 2}, {3, 4}, {1, 3}, {2, 3}},
		},
		{
			args: args{
				n: 5,
				k: 3,
			},
			want: [][]int{{1, 2, 5}, {1, 3, 5}, {2, 3, 5}, {1, 2, 3}, {1, 4, 5}, {2, 4, 5}, {1, 2, 4}, {3, 4, 5}, {1, 3, 4}, {2, 3, 4}},
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("case-%d", i), func(t *testing.T) {
			if got := Combk(tt.args.n, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("combk() = %v, want %v", got, tt.want)
			}
		})
	}
}
