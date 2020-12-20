package main

import (
	"fmt"
	"reflect"
	"testing"
)

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

func Test_allConstruct_all(t *testing.T) {
	type args struct {
		target   string
		elements []string
	}
	tests := []struct {
		name string
		c    allConstruct
		args args
		want [][]string
	}{
		{
			c: make(allConstruct),
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
			c: make(allConstruct),
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
		{
			c: make(allConstruct),
			args: args{
				target:   "eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeef",
				elements: []string{"e", "ee", "eeee", "eeeeee", "eeeeeeeeee"},
			},
			want: nil,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("case-%d", i), func(t *testing.T) {
			if got := tt.c.all(tt.args.target, tt.args.elements); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("allConstruct.all() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_countConstructTabulation(t *testing.T) {
	type args struct {
		target   string
		elements []string
	}
	tests := []struct {
		name string
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
			if got := countConstructTabulation(tt.args.target, tt.args.elements); got != tt.want {
				t.Errorf("countConstructTabulation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_allConstructTabulation(t *testing.T) {
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
			if got := allConstructTabulation(tt.args.target, tt.args.elements); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("allConstructTabulation() = %v, want %v", got, tt.want)
			}
		})
	}
}
