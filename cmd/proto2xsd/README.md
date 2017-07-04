# proto2xsd

XSD convertion tool for Google ProtocolBuffers version 3

	> proto2xsd -help
		Usage of proto2xsd [flags] [path ...]
  		-w	write result to XSD files instead of stdout

## Docker
A Docker image is available on Dockerhub.
It can be used as part of your continuous integration build pipeline.

### build 
	GOOS=linux go build
	docker build -t emicklei/proto2xsd .

### run
	 docker run emicklei/proto2xsd

© 2017, [ernestmicklei.com](http://ernestmicklei.com).  MIT License.     