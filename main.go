package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("我是——密钥")
	secret := peepSecret("secret.txt")
	code := readCode("main.go")
	fmt.Println(code.content)
	Report(scan(secret, code))
}

type File struct {
	path string
	content []string
}

// TODO: Recursively scan files under the current directory

func peepSecret(path string) *File {
	c, err := ioutil.ReadFile(path)
	if err != nil {
		panic(fmt.Sprintf("Read Secret failed: err[%s]\n", err))
	}
	// TODO: Different separators for different os
	return &File{path, strings.Split(string(c), "\r\n")}
}

func readCode(path string) *File {
	c, err := ioutil.ReadFile(path)
	if err != nil {
		panic(fmt.Sprintf("Read code failed: err[%s]\n", err))
	}
	return &File{path, strings.Split(string(c), "\n")}
}

func scan(s, c *File) [][]string {
	report := make([][]string, 0)
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
	return report
}

func Report(report [][]string) {
	if len(report) == 0 {
		fmt.Print("No secret found, your code is safe to release!")
	} else {
		fmt.Println("Warning! Secrets found!")
		fmt.Println("Secret Report: ")
		fmt.Println("File path   |   Position   |   Secret")
		for _, rep := range report {
			fmt.Printf("%s | %s | %s\n", rep[0], rep[1], rep[2])
		}
	}
}