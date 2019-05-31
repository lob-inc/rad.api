module github.com/lob-inc/rad.api

go 1.12

require (
	github.com/go-openapi/loads v0.18.0
	github.com/lob-inc/rssp/server v0.1.0
	github.com/spf13/pflag v1.0.3
)

replace github.com/lob-inc/rssp/server => ./libs/rad/server
