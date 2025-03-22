package main

import (
	"fmt"
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
	err := os.Remove(exists)
	if err != nil {
		fmt.Printf("remove failed : %v", err)
	}
}
