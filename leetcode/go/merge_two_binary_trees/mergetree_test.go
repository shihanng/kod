package mergetree

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_stringToTreeNode(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		args args
		want *TreeNode
	}{
		{
			args: args{
				input: "1,3,2,5",
			},
			want: &TreeNode{
				Val:   1,
				Left:  &TreeNode{Val: 3, Left: &TreeNode{Val: 5}},
				Right: &TreeNode{Val: 2}},
		},
		{
			args: args{
				input: "2,1,3,null,4,null,7",
			},
			want: &TreeNode{
				Val:   2,
				Left:  &TreeNode{Val: 1, Right: &TreeNode{Val: 4}},
				Right: &TreeNode{Val: 3, Right: &TreeNode{Val: 7}}},
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("case-%d", i), func(t *testing.T) {
			if got := stringToTreeNode(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("stringToTreeNode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mergeTrees(t *testing.T) {
	type args struct {
		t1 *TreeNode
		t2 *TreeNode
	}
	tests := []struct {
		args args
		want *TreeNode
	}{
		{
			args: args{
				t1: stringToTreeNode("1,3,2,5"),
				t2: stringToTreeNode("2,1,3,null,4,null,7"),
			},
			want: stringToTreeNode("3,4,5,5,4,null,7"),
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("case-%d", i), func(t *testing.T) {
			if got := mergeTrees(tt.args.t1, tt.args.t2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeTrees() = %v, want %v", got, tt.want)
			}
		})
	}
}
