package countconstruct

import (
	"fmt"
	"testing"
)

func Test_countConstruct(t *testing.T) {
	type args struct {
		target   string
		elements []string
	}
	tests := []struct {
		name string
		c    memo
		args args
		want int
	}{
		{
			args: args{
				target:   "abcdef",
				elements: []string{"ab", "abc", "cd", "def", "abcd"},
			},
			want: 1,
		},
		{
			args: args{
				target:   "skateboard",
				elements: []string{"bo", "rd", "ate", "t", "ska", "sk", "boar"},
			},
			want: 0,
		},
		{
			args: args{
				target:   "purple",
				elements: []string{"purp", "p", "ur", "le", "purpl"},
			},
			want: 2,
		},
		{
			args: args{
				target:   "enterapotentpot",
				elements: []string{"a", "p", "ent", "enter", "ot", "o", "t"},
			},
			want: 4,
		},
		{
			args: args{
				target:   "eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeef",
				elements: []string{"e", "ee", "eeee", "eeeeee", "eeeeeeeeee"},
			},
			want: 0,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("case-%d", i), func(t *testing.T) {
			m := make(memo)
			if got := m.recursive(tt.args.target, tt.args.elements); got != tt.want {
				t.Errorf("memo.recursive() = %v, want %v", got, tt.want)
			}

			if got := tabulation(tt.args.target, tt.args.elements); got != tt.want {
				t.Errorf("tabulation() = %v, want %v", got, tt.want)
			}
		})
	}
}
