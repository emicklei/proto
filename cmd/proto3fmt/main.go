package main

import (
	"log"
	"os"

	"github.com/emicklei/proto3"
)

func main() {
	p := proto3.NewParser(os.Stdin)
	def, err := p.Parse()
	if err != nil {
		log.Fatalln("proto3fmt failed:", p.Line(), err)
	}
	log.Printf("%#v", def)
}
