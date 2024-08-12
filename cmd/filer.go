package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"hash"
	"io"
	"os"
)

func GenHash(filename string) (hash.Hash, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	h := sha256.New()
	if _, err := io.Copy(h, file); err != nil {
		return nil, err
	}

	return h, nil
}

func main() {
	args := os.Args
	if len(args) != 3 {
		fmt.Println("Incorrect amount of arguments")
		return
	}

	primary, err := GenHash(os.Args[1])
	if err != nil {
		fmt.Println(err)
	}

	secondary, err := GenHash(os.Args[2])
	if err != nil {
		fmt.Println(err)
	}

	primaryHash := primary.Sum(nil)
	secondaryHash := secondary.Sum(nil)

	if bytes.Equal(primaryHash, secondaryHash) {
		fmt.Println("Hash of both files match")
	} else {
		fmt.Printf("Hash of %s does not match %s's hash\n", args[1], args[2])
	}

	fmt.Printf("Filename: %s Hash: %x\n", args[1], primaryHash)
	fmt.Printf("Filename: %s Hash: %x\n", args[2], secondaryHash)
}
