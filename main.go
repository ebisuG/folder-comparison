package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	//recieve arguments from stdin
	cli := &CLI{args: os.Args[1:]} // os.Args[0] はプログラムの実行パスなので除外
	filePaths, err := cli.receiveArguments()
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	fmt.Println("Received file paths:", filePaths)
	//check format of arguments
	//run logic

	//there is two arguments for separate file
	//repeat loop to recursively search files under a folder

	//in a loop, calc hash of a file
	//and add that value to the accumulated value
	//at the end, return the sum of hash value

	//finally, compare those values if they are same or not

}

const (
	OK    = 0
	Error = 1
)

type CLI struct {
	args []string
}

func (c *CLI) receiveArguments() ([]string, error) {

	if len(c.args) < 2 {
		return nil, errors.New("need at least 2 arguments")
	}

	for _, v := range c.args {
		if _, err := os.Stat(v); err != nil {
			return nil, fmt.Errorf("no such file: %v", v)
		}
	}

	return c.args, nil
}

func getFilePath() {

}

func searchAndApplyFunction() {

}

func calculateHash() {

}

func accumulateValue() {

}

// apply a function for some files in a parallel way
func parallelize() {

}

func doesExist(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	} else if errors.Is(err, os.ErrNotExist) {
		return false, nil
	} else {
		return false, errors.New("unknown error")
	}
}
