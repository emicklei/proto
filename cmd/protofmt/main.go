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
		if err := format(each, os.Stdout); err != nil {
			println(each, err.Error())
		}
	}
}

func format(input string, output io.Writer) error {
	content, err := ioutil.ReadFile(input)
	if err != nil {
		return err
	}
	p := proto.NewParser(bytes.NewReader(content))
	def, err := p.Parse()
	if err != nil {
		return err
	}
	proto.NewFormatter(output, "\t").Format(def)
	return nil
}
