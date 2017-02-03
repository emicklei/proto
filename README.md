# proto

[![Build Status](https://travis-ci.org/emicklei/proto.png)](https://travis-ci.org/emicklei/proto)
[![Go Report Card](https://goreportcard.com/badge/github.com/emicklei/proto)](https://goreportcard.com/report/github.com/emicklei/proto)
[![GoDoc](https://godoc.org/github.com/emicklei/proto?status.svg)](https://godoc.org/github.com/emicklei/proto)

Package in Go for parsing and formatting Google Protocol Buffers [.proto files version 2 + 3] (https://developers.google.com/protocol-buffers/docs/reference/proto-spec)

### usage

    parser := proto.NewParser(anIOReader)
	proto, err := parser.Parse()
	if err != nil {
		log.Fatalln("proto parsing failed", err)
	}
	formatter := proto.NewFormatter(anIOWriter," ")
	formatter.Format(proto)

### install

    go get -u -v github.com/emicklei/proto

© 2017, ernestmicklei.com.  MIT License. Contributions welcome.