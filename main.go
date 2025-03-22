package main

import (
	"errors"
	"os"
)

func main() {
	//recieve arguments from stdin
	//check format of arguments
	//run logic

	//there is two arguments for separate file
	//repeat loop to recursively search files under a folder

	//in a loop, calc hash of a file
	//and add that value to the accumulated value
	//at the end, return the sum of hash value

	//finally, compare those values if they are same or not

}

func receiveArguments() {

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
