package main

import "testing"

func Test_nps(t *testing.T) {
	/*scores := []int{10, 7, 10, 10, 10}
	want := 80
	got := nps(scores)
	if  want!=got{
		t.Error("nps with args:", scores,"want:",want,"got:",got)
	}*/
	//func расчет Расстояния (2 парметра) 1{} - 1 test
	//func bonus 2 tests
	/*bonus([]int)int{

	}*/
	testCases := []struct{
		name string
		scores []int
		want int
	}{
		{"All promoters",[]int{10,10,10},100},
		{"All detractors",[]int{0,0,0},-100},
		{"All neutrals",[]int{8,8,8},0},
		{"All mixed",[]int{10,10,8},66},

	}
	for _,testCase :=range testCases{
		got :=nps(testCase.scores)
		if got != testCase.want{
			t.Error("nps test:",testCase.name,"got:",got,"want:",testCase.want)
		}
	}
}