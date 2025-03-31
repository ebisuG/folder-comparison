package main

import (
	"fmt"
	"os"
	"reflect"
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

func Test_calculateHash(t *testing.T) {
	testFile := "toBeCalculateHash.txt"
	tests := []struct {
		name      string
		filePath  string
		expected  []byte
		wantError bool
	}{
		{"caluculate hash", testFile, []byte{45, 33, 42, 30, 43, 34, 34, 32, 39, 38, 46, 43, 31, 43, 31, 34, 39, 41, 46, 42, 46, 34, 43, 38, 39, 39, 36, 46, 42, 39, 32, 34, 32, 37, 41, 45, 34, 31, 45, 34, 36, 34, 39, 42, 39, 33, 34, 43, 41, 34, 39, 35, 39, 39, 31, 42, 37, 38, 35, 32, 42, 38, 35, 35}, false},
		{"failed to calculate", "nofile.txt", []byte{0}, true},
	}
	os.Create(testFile)
	defer os.RemoveAll(testFile)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hash, err := calculateHash(tt.filePath)
			if tt.wantError {
				if err == nil {
					fmt.Errorf("should be err : %v", tt.name)
				}
				if hash != nil {
					fmt.Errorf("should be empty string : %v", tt.name)
					fmt.Errorf("return : %v", hash)
				}
			} else {
				if reflect.DeepEqual(hash, tt.expected) {
					fmt.Errorf("test case : %v", tt.name)
					fmt.Errorf("In %v, hash is wrong. \n expected : %v \n result : %v", tt.filePath, tt.expected, hash)
				}
			}
		})
	}

}
