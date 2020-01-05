package main

import "testing"

func Test_nps(t *testing.T) {

	scores := []int{10, 7, 10, 10, 10}
	want := 80
	got := nps(scores)
	if  want!=got{
		t.Error("nps with args:", scores,"want:",want,"got:",got)
	}
}