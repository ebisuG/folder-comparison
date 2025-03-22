package main

import (
	"os"
	"testing"
)

func Test_doesExist(t *testing.T) {
	exists := "test.txt"
	notExist := "noSuchFile.txt"
	os.Create("test.txt")
	expectedTrue, err1 := doesExist(exists)
	expectedFalse, err2 := doesExist(notExist)
	if !(expectedTrue && !expectedFalse) {
		t.Errorf(`doesExist doesn't work well`)
		t.Errorf(`exists : %v`, err1)
		t.Errorf(`notExist : %v`, err2)
	}
	os.Remove(exists)
}
