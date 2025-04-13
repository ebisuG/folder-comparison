package main

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sync"
)

func main() {
	//recieve arguments from stdin
	cli := &CLI{args: os.Args[1:]} // os.Args[0] はプログラムの実行パスなので除外
	filePaths, err := cli.receiveArguments()
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	fmt.Println("Compare :", filePaths)
	fmt.Println("==========")

	var wg sync.WaitGroup
	results := make([]string, len(cli.args))

	for i, v := range cli.args {
		wg.Add(1)
		go func(i int, v string) {
			defer wg.Done()
			fh := &FileHash{hash: []byte{}, rootFolder: v}
			hash, err := fh.CalcHashRecursively()
			if err != nil {
				fmt.Println("Error:", err)
				os.Exit(1)
			}
			results[i] = hash
			fmt.Printf("%v : %v \n", hash, v)
		}(i, v)
	}

	wg.Wait()

	for i := range results {
		if i != len(results)-1 {
			if results[i] == results[i+1] {
				continue
			} else {
				fmt.Println("==========")
				fmt.Println("Different files")
				return
			}
		}
	}
	fmt.Println("==========")
	fmt.Printf("Same files")

}

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

type FileHash struct {
	hash       []byte
	rootFolder string
}

func (fh *FileHash) CalcHashRecursively() (string, error) {
	err := filepath.WalkDir(
		fh.rootFolder,
		func(path string, d os.DirEntry, err error) error {
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

	return h.Sum(nil), nil

}
