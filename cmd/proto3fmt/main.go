package main

import (
	"log"
	"os"

	"github.com/emicklei/proto3"
)

// go run *.go < example1.proto
// go run *.go < example0.proto
func main() {
	p := proto3.NewParser(os.Stdin)
	def, err := p.Parse()
	if err != nil {
		log.Fatalln("proto3fmt failed, on line", p.Line(), err)
	}
	f := &formatter{w: os.Stdout, indentSeparator: "    "}
	for _, each := range def.Elements {
		each.Accept(f)
	}
}
