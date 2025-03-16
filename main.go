package main

import (
	"errors"
	"os"
)

func main() {

}

func getFilePath() {

}

func doesExist(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	} else if errors.Is(err, os.ErrNotExist) {
		return false, nil
	} else {
		return false, errors.New("Unknown Error")
	}
}
