package canconstruct

import (
	"fmt"
	"testing"
)

func Test_canconstruct(t *testing.T) {
	type args struct {
		target   string
		elements []string
	}
	tests := []struct {
		args args
		want bool
	}{
		{
			args: args{
				target:   "abcdef",
				elements: []string{"ab", "abc", "cd", "def", "abcd"},
			},
			want: true,
		},
		{
			args: args{
				target:   "skateboard",
				elements: []string{"bo", "rd", "ate", "t", "ska", "sk", "boar"},
			},
			want: false,
		},
		{
			args: args{
				target:   "eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeef",
				elements: []string{"e", "ee", "eeee", "eeeeee", "eeeeeeeeee"},
			},
			want: false,
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
