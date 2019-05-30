package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

var report = make([][]string, 0)

func main() {
	secretPath := flag.String("s", "", "secret filepath")
	separator := flag.String("sep", "", "separator")

	flag.Parse()

	if *secretPath == "" {
		fmt.Printf("Please indicate a path for secret!\n  Example: NoSecretLeak -s=secret\n")
		return
	}

	if *separator == "" {
		fmt.Printf("Please indicate a separator for secret!\n  Example: NoSecretLeak -sep=,\n")
		return
	}

	path, err := os.Getwd()
	if err != nil {
		fmt.Printf("Get current working directory failed: [%s]\n", err)
		return
	}

	secret, err := PeepSecret(*secretPath)
	if err != nil {
		fmt.Printf("Get secret failed: [%s]\n", err)
		return
	}

	err = filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			code, errRC := ReadCode(path)
			if errRC != nil {
				return errRC
			}
			Scan(secret, code)
		}
		return nil
	})
	if err != nil {
		fmt.Printf("Error Encountered during walking: [%s]\n", err)
	}

	Report(report)
}

type File struct {
	path string
	content []string
}

// TODO: Recursively scan files under the current directory

func PeepSecret(path string) (*File, error) {
	c, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	//fmt.Printf("byte secret: [%v]", c)
	// TODO: Different separators for different os
	// for .txt file, use:
	// 		\n for mac
	//	  \r\n for windows
	return &File{path, strings.Split(string(c), ",")}, nil
}

func ReadCode(path string) (*File, error) {
	c, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return &File{path, strings.Split(string(c), "\n")}, nil
}

func Scan(s, c *File) {
	secret := s.content
	code := c.content
	for i, line := 0, 1; i < len(code); i++{
		for j, column := 0, 1; j < len(code[i]); j++ {
			if(code[i][j] == 9) {
				column += 3
			}
			for k := 0; k < len(secret); k++ {
				l := len(secret[k])
				if( j+l <= len(code[i]) && strings.EqualFold(code[i][j:j+l], secret[k]) ) {
					report = append(report,
						[]string{c.path, strconv.Itoa(line)+":"+strconv.Itoa(column), secret[k]})
				}
			}
			column++
		}
		line++
	}
}

func Report(report [][]string) {
	if len(report) == 0 {
		fmt.Print("No secret found, your code is safe to release!\n")
	} else {
		fmt.Println("Warning! Secrets found!")
		fmt.Println("Secret Report: ")
		fmt.Println("File path   |   Position   |   Secret")
		for _, rep := range report {
			fmt.Printf("%s | %s | %s\n", rep[0], rep[1], rep[2])
		}
	}
}