# proto

[![Build Status](https://travis-ci.org/emicklei/proto.png)](https://travis-ci.org/emicklei/proto)
[![GoDoc](https://godoc.org/github.com/emicklei/proto?status.svg)](https://godoc.org/github.com/emicklei/proto)

Package in Go for parsing Google Protocol Buffers [.proto files version 3] (https://developers.google.com/protocol-buffers/docs/reference/proto-spec)

### usage

    parser := proto.NewParser(anIOReader)
	proto, err := parser.Parse()
	if err != nil {
		log.Fatalln("proto parsing failed", err)
	}

### install

    go get -u -v github.com/emicklei/proto

© 2017, ernestmicklei.com.  MIT License. Contributions welcome.