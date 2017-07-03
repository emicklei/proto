# protofmt

formatting tool for Google ProtocolBuffers version 2 and 3

	> protofmt -help
		Usage of protofmt [flags] [path ...]
  		-w	write result to (source) file instead of stdout

## Docker
A Docker image is available on Dockerhub.
It can be used as part of your continuous integration build pipeline.

### build 
	GOOS=linux go build
	docker build -t emicklei/protofmt .

### run
	 docker run emicklei/protofmt	

© 2017, [ernestmicklei.com](http://ernestmicklei.com).  MIT License.     