package main

import (
	"fmt"
	"reflect"
	"testing"
)

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
