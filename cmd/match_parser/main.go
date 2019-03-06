package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/davecgh/go-spew/spew"
	"github.com/dvilchansky/gopubg/models/match"
	"github.com/sirupsen/logrus"
)

var (
	input string
)

func usage() {
	fmt.Printf("usage: %s [options]\n", os.Args[0])
	flag.PrintDefaults()
	os.Exit(1)
}

func init() {
	// Parameters
	flag.StringVar(&input, "input", "", "input file")

	// Parse parameters
	flag.Parse()

	// Verify parameters
	if input == "" {
		usage()
	}
}

func main() {
	file, err := os.Open(input)
	if err != nil {
		logrus.Fatal(err)
	}

	m, err := match.ParseMatch(file)
	if err != nil {
		logrus.Fatal(err)
	}

	spew.Dump(m)
}
