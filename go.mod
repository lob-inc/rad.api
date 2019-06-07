module github.rakops.com/gatd/rad.api

go 1.12

require (
	github.com/go-openapi/loads v0.18.0
	github.com/grpc-ecosystem/grpc-gateway v1.9.0
	github.com/lob-inc/rssp/server v0.1.0
	github.com/spf13/pflag v1.0.3
	github.rakops.com/gatd/rad.api.pb v0.0.0-20190607041523-c2a2a624e2ce
	google.golang.org/grpc v1.21.0
)

replace github.com/lob-inc/rssp/server => ./libs/rad/server
