# proto

[![Build Status](https://travis-ci.org/emicklei/proto.png)](https://travis-ci.org/emicklei/proto)
[![Go Report Card](https://goreportcard.com/badge/github.com/emicklei/proto)](https://goreportcard.com/report/github.com/emicklei/proto)
[![GoDoc](https://godoc.org/github.com/emicklei/proto?status.svg)](https://godoc.org/github.com/emicklei/proto)

Package in Go for parsing Google Protocol Buffers [.proto files version 2 + 3] (https://developers.google.com/protocol-buffers/docs/reference/proto3-spec)

This repository also includes 2 commands. The `protofmt` tool is for formatting .proto files and the `proto2xsd` tool is for generating XSD files from .proto version 3 files.

### usage as package

    parser := proto.NewParser(reader)
	definition, err := parser.Parse()
	if err != nil {
		log.Fatalln("proto parsing failed", err)
	}

	formatter := proto.NewFormatter(writer," ")
	formatter.Format(definition)

### usage of proto2xsd command

	> proto2xsd -help
		Usage of proto2xsd [flags] [path ...]
  		-w	write result to an XSD files instead of stdout

See folder `cmd/proto2xsd/README.md` for more details.

### usage of protofmt command

	> protofmt -help
		Usage of protofmt [flags] [path ...]
  		-w	write result to (source) files instead of stdout

See folder `cmd/protofmt/README.md` for more details.

### how to install

    go get -u -v github.com/emicklei/proto

© 2017, [ernestmicklei.com](http://ernestmicklei.com).  MIT License. Contributions welcome.