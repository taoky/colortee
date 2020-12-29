package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
)

var colorRe = regexp.MustCompile(`\x1b[^m]*m`)

func strip(str string) string {
	return colorRe.ReplaceAllString(str, "")
}

func tee(in io.ReadSeeker, consoleOuts []io.Writer, fileOuts []io.Writer) error {
	scanner := bufio.NewScanner(in)

	for scanner.Scan() {
		text := scanner.Text()
		strippedText := strip(text)
		for _, out := range consoleOuts {
			_, err := fmt.Fprintln(out, text)
			if err != nil {
				return err
			}
		}
		for _, out := range fileOuts {
			_, err := fmt.Fprintln(out, strippedText)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func main() {
	appendPtr := flag.Bool("append", false, "append rather than overwrite to file is set.")
	rawPtr := flag.Bool("raw", false, "output unprocessed raw to file")

	flag.Parse()

	if len(flag.Args()) != 1 {
		log.Fatalf("Should have exactly one output file. Example: %s output.txt", os.Args[0])
	}

	outputFileName := flag.Args()[0]

	fileFlag := os.O_APPEND
	if !*appendPtr {
		fileFlag = os.O_TRUNC
	}
	outputFile, err := os.OpenFile(outputFileName, os.O_CREATE|os.O_WRONLY|fileFlag, 0644)
	if err != nil {
		log.Fatal(err)
	}

	var consoleOuts []io.Writer
	var fileOuts []io.Writer

	consoleOuts = append(consoleOuts, os.Stdout)
	if *rawPtr {
		consoleOuts = append(consoleOuts, outputFile)
	} else {
		fileOuts = append(fileOuts, outputFile)
	}

	err = tee(os.Stdin, consoleOuts, fileOuts)
	if err != nil {
		log.Fatal(err)
	}
}
