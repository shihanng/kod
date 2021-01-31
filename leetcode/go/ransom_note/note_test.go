package note

import (
	"fmt"
	"testing"
)

func Test_canConstruct(t *testing.T) {
	type args struct {
		ransomNote string
		magazine   string
	}
	tests := []struct {
		args args
		want bool
	}{
		{
			args: args{
				ransomNote: "a",
				magazine:   "b",
			},
			want: false,
		},
		{
			args: args{
				ransomNote: "aa",
				magazine:   "ab",
			},
			want: false,
		},
		{
			args: args{
				ransomNote: "aa",
				magazine:   "aab",
			},
			want: true,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("case-%d", i), func(t *testing.T) {
			if got := canConstruct(tt.args.ransomNote, tt.args.magazine); got != tt.want {
				t.Errorf("canConstruct() = %v, want %v", got, tt.want)
			}
		})
	}
}
