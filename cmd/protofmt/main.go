package main

import (
	"log"
	"os"

	"github.com/emicklei/proto"
)

// go run *.go unformatted.proto
func main() {
	if len(os.Args) == 1 {
		log.Fatal("Usage: protofmt my.proto")
	}
	i, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer i.Close()
	p := proto.NewParser(i)
	def, err := p.Parse()
	if err != nil {
		log.Fatalln("protofmt failed", err)
	}
	proto.NewFormatter(os.Stdout, "  ").Format(def)
	//spew.Dump(def)
}
