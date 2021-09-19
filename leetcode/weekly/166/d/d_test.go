// Code generated by generator_test.
package main

import (
	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"testing"
)

func Test(t *testing.T) {
	t.Log("Current test is [d]")
	exampleIns := [][]string{{`[[0,0],[0,1]]`}, {`[[0]]`}, {`[[1,1,1],[1,0,1],[0,0,0]]`}, {`[[1,0,0],[1,0,0]]`}}
	exampleOuts := [][]string{{`3`}, {`0`}, {`6`}, {`-1`}}
	// custom test cases or WA cases.
	//exampleIns = append(exampleIns, []string{``})
	//exampleOuts = append(exampleOuts, []string{``})
	if err := testutil.RunLeetCodeFunc(t, minFlips, exampleIns, exampleOuts); err != nil {
		t.Fatal(err)
	}
}