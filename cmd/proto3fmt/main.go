package main

import (
	"log"
	"os"

	"github.com/emicklei/proto3"
)

// go run *.go example1.proto
// go run *.go example0.proto
func main() {
	if len(os.Args) == 1 {
		log.Fatal("Usage: proto3fmt my.proto")
	}
	i, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer i.Close()
	p := proto3.NewParser(i)
	def, err := p.Parse()
	if err != nil {
		log.Fatalln("proto3fmt failed, on line", p.Line(), err)
	}
	f := &formatter{w: os.Stdout, indentSeparator: "    "}
	for _, each := range def.Elements {
		each.Accept(f)
	}
}
