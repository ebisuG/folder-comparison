package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
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
	os.Mkdir("./folder1-1", 0777)
	os.Mkdir("folder1-2", 0777)
	defer os.RemoveAll("./folder1-1/")
	defer os.RemoveAll("./folder1-2/")
	correctInput1 := "./folder1-1 folder1-2"
	wrongInput1 := ""
	wrongInput2 := " ./wrongFile2-1"
	wrongTestInputs := []string{wrongInput1, wrongInput2}
	outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)
	cli := &CLI{outStream: outStream, errStream: errStream}

	args := strings.Split(correctInput1, " ")
	fmt.Println(args)
	result, err := cli.receiveArguments(args)
	if err != nil {
		t.Errorf("return value should be nil : %v \n", err)
	}
	if !(result[0] == "./folder1-1" && result[1] == "folder1-2") {
		t.Errorf("length of return value should be two : %v \n", err)
	}

	for _, v := range wrongTestInputs {
		args := strings.Split(v, " ")
		result, err := cli.receiveArguments(args)
		if err == nil {
			t.Errorf("receiveArguments should fail : %v \n", err)
		}
		if len(result) != 0 {
			t.Errorf("receiveArguments should return empty array : %v \n", err)
		}
	}

}
