package main

import (
	"fmt"
	"log"
	"os"

	loads "github.com/go-openapi/loads"

	flag "github.com/spf13/pflag"

	"github.com/lob-inc/rssp/server"
	"github.com/lob-inc/rssp/server/api/restapi"
	"github.com/lob-inc/rssp/server/api/restapi/operations"
)

func main() {
	swaggerSpec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
	if err != nil {
		log.Fatalln(err)
	}

	var srv *restapi.Server // make sure init is called

	flag.Usage = func() {
		fmt.Fprint(os.Stderr, "Usage:\n")
		fmt.Fprint(os.Stderr, "  rssp-server [OPTIONS]\n\n")

		title := "RSSP"
		fmt.Fprint(os.Stderr, title+"\n\n")
		desc := "API for UI\n"
		if desc != "" {
			fmt.Fprintf(os.Stderr, desc+"\n\n")
		}
		fmt.Fprintln(os.Stderr, flag.CommandLine.FlagUsages())
	}
	// parse the CLI flags
	flag.Parse()

	api := operations.NewRsspAPI(swaggerSpec)
	// get server with flag values filled out
	srv = restapi.NewServer(api)

	conf := server.NewConfig()
	restapi.InjectServiceDependencies(api, conf)

	operations.ConfigureServices(api)
	defer srv.Shutdown()
	srv.ConfigureAPI()

	if srv.Port == 0 {
		srv.Port = conf.Server.APIPort
	}

	if err := srv.Serve(); err != nil {
		log.Fatalln(err)
	}

}
