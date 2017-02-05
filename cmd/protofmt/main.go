package main

import (
	"io"
	"os"

	"flag"

	"bytes"
	"io/ioutil"

	"github.com/emicklei/proto"
)

var (
	overwrite = flag.Bool("w", false, "write result to (source) file instead of stdout")
)

// go run *.go unformatted.proto
func main() {
	flag.Parse()
	if len(flag.Args()) == 0 {
		flag.Usage()
		os.Exit(0)
	}
	for _, each := range flag.Args() {
		if err := readFormatWrite(each); err != nil {
			println(each, err.Error())
		}
	}
}

func readFormatWrite(filename string) error {
	// open for read
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	// buffer before write
	buf := new(bytes.Buffer)
	if err := format(file, buf); err != nil {
		return err
	}
	if *overwrite {
		// write back to input
		if err := ioutil.WriteFile(filename, buf.Bytes(), os.ModePerm); err != nil {
			return err
		}
	} else {
		// write to stdout
		if _, err := io.Copy(os.Stdout, bytes.NewReader(buf.Bytes())); err != nil {
			return err
		}
	}
	return nil
}

func format(input io.Reader, output io.Writer) error {
	parser := proto.NewParser(input)
	def, err := parser.Parse()
	if err != nil {
		return err
	}
	proto.NewFormatter(output, "\t").Format(def)
	return nil
}
