package main

import (
	"flag"
	"fmt"
	//"github.com/pkg/errors"
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
	del := flag.Bool("d", false, "indicate whether to delete the secret file after scanning")

	flag.Parse()

	if *secretPath == "" {
		fmt.Printf("Please indicate a path for secret!\n  Example: NoSecretLeak -s=secret\n")
		os.Exit(4)
	}

	if *separator == "" {
		fmt.Printf("Please indicate a separator for secret!\n  Example: NoSecretLeak -sep=,\n")
		os.Exit(4)
	}

	path, err := os.Getwd()
	if err != nil {
		fmt.Printf("Get current working directory failed: [%s]\n", err)
		os.Exit(4)
	}

	// Extract wordlist from secret file
	secret, err := PeepSecret(*secretPath, *separator)
	if err != nil {
		fmt.Printf("Get secret failed: filepath: [%s] err: [%s]\n", *secretPath, err)
		os.Exit(4)
	}

	// Recursively traverse all files under current working directory
	err = filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// secret file is skipped
		if !info.IsDir() && info.Name() != secret.path {
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
		fmt.Printf("Files under current directories may not be fully scaned, your are recommended to scan secret again\n")
	}

	// Print report to user
	Report(report)

	if *del {
		errRemoveSecret := os.Remove(*secretPath)
		if errRemoveSecret != nil {
			fmt.Printf("Secret file auto-delete failed: [%s]\n", errRemoveSecret)
			fmt.Printf("Warning! Secret file auto-delete failed, plesat delete it manully!")
			fmt.Printf("Warning! Secret file auto-delete failed, plesat delete it manully!")
			fmt.Printf("Warning! Secret file auto-delete failed, plesat delete it manully!")
		} else {
			fmt.Println("Secret file is deleted!")
		}
	}

	// Secret found, using exit code 3 to indicate bash script
	if len(report) > 0 {
		os.Exit(3)
	}
}

type File struct {
	path string
	content []string
}

func PeepSecret(path, sep string) (*File, error) {
	c, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	raw := strings.Split(string(c[:len(c)-1]), sep)
	content := make([]string, 0)
	for i := 0; i<len(raw); i++ {
		if raw[i] == "" {
			continue
		}
		content = append(content, raw[i])
	}
	return &File{path, content}, nil
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
			if code[i][j] == 9 {
				column += 3
			}
			for k := 0; k < len(secret); k++ {
				l := len(secret[k])
				if j+l <= len(code[i]) && strings.EqualFold(code[i][j:j+l], secret[k])  {
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