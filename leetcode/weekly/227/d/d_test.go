// Code generated by copypasta/template/leetcode/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"testing"
)

func Test(t *testing.T) {
	t.Log("Current test is [d]")
	examples := [][]string{
		{
			`[5,-7,3,5]`, `6`, 
			`0`,
		},
		{
			`[7,-9,15,-2]`, `-5`, 
			`1`,
		},
		{
			`[1,2,3]`, `-7`, 
			`7`,
		},
		// TODO 测试入参最小的情况
		{
			`[-2772,6927,4773,-2687,7167,-8995,2940,8869,526]`, `969621127`,
			`969589925`,
		},
		{
			`[200,100]`, `1`,
			`1`,
		},
	}
	targetCaseNum :=  -1
	if err := testutil.RunLeetCodeFuncWithExamples(t, minAbsDifference, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
// https://leetcode-cn.com/contest/weekly-contest-227/problems/closest-subsequence-sum/