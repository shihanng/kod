package main

import (
	"fmt"
	"testing"
)

func Test_fibTabulation(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		args args
		want int
	}{
		{
			args: args{
				n: 6,
			},
			want: 8,
		},
		{
			args: args{
				n: 7,
			},
			want: 13,
		},
		{
			args: args{
				n: 50,
			},
			want: 12586269025,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("case-%d", i), func(t *testing.T) {
			if got := fibTabulation(tt.args.n); got != tt.want {
				t.Errorf("fibTabulation() = %v, want %v", got, tt.want)
			}
		})
	}
}
