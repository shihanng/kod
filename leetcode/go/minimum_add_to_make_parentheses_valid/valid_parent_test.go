package valid_parent

import "testing"

func Test_minAddToMakeValid(t *testing.T) {
	type args struct {
		S string
	}
	tests := []struct {
		args args
		want int
	}{
		{
			args: args{
				S: "())",
			},
			want: 1,
		},
		{
			args: args{
				S: "(((",
			},
			want: 3,
		},
		{
			args: args{
				S: "()",
			},
			want: 0,
		},
		{
			args: args{
				S: "()))((",
			},
			want: 4,
		},
		{
			args: args{
				S: "((())",
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.args.S, func(t *testing.T) {
			if got := minAddToMakeValid(tt.args.S); got != tt.want {
				t.Errorf("minAddToMakeValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
