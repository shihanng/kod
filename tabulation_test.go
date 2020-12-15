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

func Test_gridTravellerTabulation(t *testing.T) {
	type args struct {
		m int
		n int
	}
	tests := []struct {
		args args
		want int
	}{
		{
			args: args{
				m: 3,
				n: 3,
			},
			want: 6,
		},
		{
			args: args{
				m: 2,
				n: 3,
			},
			want: 3,
		},
		{
			args: args{
				m: 18,
				n: 18,
			},
			want: 2333606220,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("case-%d", i), func(t *testing.T) {
			if got := gridTravellerTabulation(tt.args.m, tt.args.n); got != tt.want {
				t.Errorf("gridTravellerTabulation() = %v, want %v", got, tt.want)
			}
		})
	}
}
