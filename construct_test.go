package main

import (
	"fmt"
	"testing"
)

func Test_canConstruct_can(t *testing.T) {
	type args struct {
		target   string
		elements []string
	}
	tests := []struct {
		c    canConstruct
		args args
		want bool
	}{
		{
			c: make(canConstruct),
			args: args{
				target:   "abcdef",
				elements: []string{"ab", "abc", "cd", "def", "abcd"},
			},
			want: true,
		},
		{
			c: make(canConstruct),
			args: args{
				target:   "skateboard",
				elements: []string{"bo", "rd", "ate", "t", "ska", "sk", "boar"},
			},
			want: false,
		},
		{
			c: make(canConstruct),
			args: args{
				target:   "eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeef",
				elements: []string{"e", "ee", "eeee", "eeeeee", "eeeeeeeeee"},
			},
			want: false,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("case-%d", i), func(t *testing.T) {
			if got := tt.c.can(tt.args.target, tt.args.elements); got != tt.want {
				t.Errorf("canConstruct.can() = %v, want %v", got, tt.want)
			}
		})
	}
}
