package allconstruct

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_allconstruct(t *testing.T) {
	type args struct {
		target   string
		elements []string
	}
	tests := []struct {
		args args
		want [][]string
	}{
		{
			args: args{
				target:   "purple",
				elements: []string{"purp", "p", "ur", "le", "purpl"},
			},
			want: [][]string{
				{"purp", "le"},
				{"p", "ur", "p", "le"},
			},
		},
		{
			args: args{
				target:   "abcdef",
				elements: []string{"ab", "abc", "cd", "def", "abcd", "ef", "c"},
			},
			want: [][]string{
				{"abc", "def"},
				{"ab", "c", "def"},
				{"abcd", "ef"},
				{"ab", "cd", "ef"},
			},
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("case-%d", i), func(t *testing.T) {
			if got := tabulation(tt.args.target, tt.args.elements); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("tabulation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_allconstruct_fast(t *testing.T) {
	type args struct {
		target   string
		elements []string
	}
	tests := []struct {
		args args
		want [][]string
	}{
		{
			args: args{
				target:   "eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeef",
				elements: []string{"e", "ee", "eeee", "eeeeee", "eeeeeeeeee"},
			},
			want: nil,
		},
		{
			args: args{
				target:   "purple",
				elements: []string{"purp", "p", "ur", "le", "purpl"},
			},
			want: [][]string{
				{"purp", "le"},
				{"p", "ur", "p", "le"},
			},
		},
		{
			args: args{
				target:   "abcdef",
				elements: []string{"ab", "abc", "cd", "def", "abcd", "ef", "c"},
			},
			want: [][]string{
				{"ab", "cd", "ef"},
				{"ab", "c", "def"},
				{"abc", "def"},
				{"abcd", "ef"},
			},
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("case-%d", i), func(t *testing.T) {
			m := make(memo)
			if got := m.recursive(tt.args.target, tt.args.elements); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("memo.recursive() = %v, want %v", got, tt.want)
			}
		})
	}
}
