package main

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
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

type FileHash struct {
	hash       []byte
	rootFolder string
}

func (fh *FileHash) CalcHashRecursively() (string, error) {
	err := filepath.WalkDir(
		fh.rootFolder,
		func(path string, d os.DirEntry, err error) error {
			// var hashSum []byte
			if d.IsDir() {
				return nil
			} else {
				hash, err1 := calculateHash(path)
				if err1 != nil {
					return fmt.Errorf("failed to calculate hash : %v", err1)
				}
				fh.hash = append(fh.hash, hash...)
				return nil
			}
		})
	if err != nil {
		return "", err
	}
	fileHash := sha256.New()
	fileHash.Write(fh.hash)
	return hex.EncodeToString(fileHash.Sum(nil)), nil
}

func calculateHash(path string) ([]byte, error) {
	f, err := os.Open(path)
	defer f.Close()
	if err != nil {
		return []byte{0}, fmt.Errorf("no such file to calculate hash : %v", err)
	}

	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		return []byte{0}, fmt.Errorf("failed to calculate hash : %v", err)
	}

	fmt.Printf("%x", h.Sum(nil))
	return h.Sum(nil), nil

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
