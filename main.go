package main

import (
	"fmt"
	"github.com/jessevdk/go-flags"
	"io/ioutil"
	"os"
	"strings"
)

var opts struct {
	Filepath         string `short:"f" long:"file" description:"the input file" required:"true"`
	ReplaceText      string `short:"r" long:"replace" description:"the text to be replaced" required:"true"`
	SubstitutionText string `short:"s" long:"substitution" description:"the text to be substituted" required:"true"`
	PrintToStdout    bool   `short:"c" long:"to-console" description:"write to console instead of to file"`
}

func main() {
	_, err := flags.ParseArgs(&opts, os.Args)
	if err != nil {
		os.Exit(1)
	}

	inputBytes, err := ioutil.ReadFile(opts.Filepath)
	if err != nil {
		fmt.Println("error: file path could not be read")
		os.Exit(1)
	}

	inputText := string(inputBytes)
	newText := strings.ReplaceAll(inputText, opts.ReplaceText, opts.SubstitutionText)

	if opts.PrintToStdout {
		fmt.Println(newText)
	} else {
		ioutil.WriteFile(opts.Filepath, []byte(newText), 0644)
	}
}
