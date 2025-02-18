package main

import (
	"reflect"
	"testing"
)

func TestMakeBeefSummary(t *testing.T) {
	data := "Fatback t-bone t-bone, pastrami t-bone.  pork, meatloaf jowl enim.  Bresaola t-bone."
	expected := map[string]int{
		"t-bone":   4,
		"fatback":  1,
		"pastrami": 1,
		"pork":     1,
		"meatloaf": 1,
		"jowl":     1,
		"enim":     1,
		"bresaola": 1,
	}

	result := makeBeefSummary(data)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}
