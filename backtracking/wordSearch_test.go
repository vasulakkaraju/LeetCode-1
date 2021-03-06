package backtracking_test

import (
	"testing"
	"github.com/g10guang/leetcode/backtracking"
)

func TestExist(t *testing.T) {
	suits := []struct{
		board [][]byte
		word string
		expect bool
	} {
		{board:[][]byte{{'a', 'a'}}, word:"aaa", expect:false},
	}
	for _, s := range suits {
		if r := backtracking.Exist(s.board, s.word); r != s.expect {
			t.Errorf("board=%s\n word=%v\n expect=%v\n return=%v\n", s.board, s.word, s.expect, r)
		}
	}
}
