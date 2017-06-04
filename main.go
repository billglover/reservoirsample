package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"time"
)

func main() {
	n := flag.Int64("n", 1, "number of samples to select (default 1)")
	i := flag.String("i", "", "input file (default stdin)")
	o := flag.String("o", "", "output file (default stdout)")
	flag.Parse()

	inf, err := openInput(*i)
	if err != nil {
		log.Fatalf("unable to open input file: %v", err)
	}
	defer inf.Close()

	outf, err := openInput(*o)
	if err != nil {
		log.Fatalf("unable to open output file: %v", err)
	}
	defer outf.Close()

	err = sampleLines(inf, outf, *n)
	if err != nil {
		log.Fatalf("unable to sample lines: %v", err)
	}
}

func sampleLines(ir io.Reader, ow io.Writer, n int64) error {

	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	reservoir := make([]string, n)
	sample := int64(0)

	scanner := bufio.NewScanner(ir)
	for scanner.Scan() {
		if sample < n {
			reservoir[sample] = scanner.Text()
		}

		if sample >= n {
			rv := random.Float64()
			p := float64(n) / float64(sample)
			capture := rv < p

			if capture {
				index := random.Intn(int(n))
				reservoir[index] = scanner.Text()
			}
		}

		sample++
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("unable to read line: %v", err)
	}

	for _, element := range reservoir {
		fmt.Fprintln(ow, element)
	}

	return nil
}

func openInput(p string) (*os.File, error) {
	if p == "" {
		return os.Stdin, nil
	}
	f, err := os.Open(p)
	return f, err
}

func openOutput(p string) (*os.File, error) {
	if p == "" {
		return os.Stdout, nil
	}
	f, err := os.Create(p)
	return f, err
}
