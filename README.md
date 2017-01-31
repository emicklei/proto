# proto3

[![Build Status](https://travis-ci.org/emicklei/proto3.png)](https://travis-ci.org/emicklei/proto3)
[![GoDoc](https://godoc.org/github.com/emicklei/proto3?status.svg)](https://godoc.org/github.com/emicklei/proto3)

Package in Go for parsing Google Protocol Buffers [.proto files version 3] (https://developers.google.com/protocol-buffers/docs/reference/proto3-spec)

### usage

    parser := proto3.NewParser(anIOReader)
	proto, err := parser.Parse()
	if err != nil {
		log.Fatalln("proto3 parsing failed", err)
	}

### install

    go get -u -v github.com/emicklei/proto3

© 2017, ernestmicklei.com.  MIT License. Contributions welcome.