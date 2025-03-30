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

func Test_receiveArguments(t *testing.T) {
	testDirs := []string{"./folder1-1", "./folder1-2"}
	for _, dir := range testDirs {
		if err := os.Mkdir(dir, 0777); err != nil {
			t.Fatalf("failed to create test directory: %v", err)
		}
	}
	defer func() {
		for _, dir := range testDirs {
			os.RemoveAll(dir)
		}
	}()

	tests := []struct {
		name      string
		args      []string
		wantError bool
	}{
		{"Valid paths", []string{"./folder1-1", "./folder1-2"}, false},
		{"Missing arguments", []string{""}, true},
		{"Non-existent path", []string{"./folder1-1", "./wrongFile"}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cli := &CLI{args: tt.args}
			result, err := cli.receiveArguments()
			if tt.wantError {
				if err == nil {
					t.Errorf("expected error but got none")
				}
			} else {
				if err != nil {
					t.Errorf("unexpected error: %v", err)
				}
				if len(result) != len(tt.args) {
					t.Errorf("expected %d arguments, got %d", len(tt.args), len(result))
				}
			}
		})
	}

}
