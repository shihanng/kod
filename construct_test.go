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

func Test_countConstruct_count(t *testing.T) {
	type args struct {
		target   string
		elements []string
	}
	tests := []struct {
		name string
		c    countConstruct
		args args
		want int
	}{
		{
			c: make(countConstruct),
			args: args{
				target:   "abcdef",
				elements: []string{"ab", "abc", "cd", "def", "abcd"},
			},
			want: 1,
		},
		{
			c: make(countConstruct),
			args: args{
				target:   "skateboard",
				elements: []string{"bo", "rd", "ate", "t", "ska", "sk", "boar"},
			},
			want: 0,
		},
		{
			c: make(countConstruct),
			args: args{
				target:   "purple",
				elements: []string{"purp", "p", "ur", "le", "purpl"},
			},
			want: 2,
		},
		{
			c: make(countConstruct),
			args: args{
				target:   "enterapotentpot",
				elements: []string{"a", "p", "ent", "enter", "ot", "o", "t"},
			},
			want: 4,
		},
		{
			c: make(countConstruct),
			args: args{
				target:   "eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeef",
				elements: []string{"e", "ee", "eeee", "eeeeee", "eeeeeeeeee"},
			},
			want: 0,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("case-%d", i), func(t *testing.T) {
			if got := tt.c.count(tt.args.target, tt.args.elements); got != tt.want {
				t.Errorf("countConstruct.count() = %v, want %v", got, tt.want)
			}
		})
	}
}
