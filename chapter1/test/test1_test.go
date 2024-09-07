package main

import "testing"

func TestSum(t *testing.T) {
	a := 5
	b := 6
	expected := 11
	res := sum(a, b)
	if res != expected {
		t.Errorf("Our sum function doens't work, %d+%d isn't %d\n", a, b,
			res)
	}
}

//func TestSum2(t *testing.T) {
//	a := 5
//	b := 6
//	expected := 10
//	res := sum(a, b)
//	if res != expected {
//		t.Errorf("Our sum function doens't work")
//	}
//}

func TestMultiply(t *testing.T) {
	a := 3
	b := 2
	expected := 6
	res := multiply(a, b)
	if res != expected {
		t.Errorf("Our multiply function doens't work")
	}
}
