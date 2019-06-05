package main

import (
	"fmt"
	"os"

	flag "github.com/spf13/pflag"

	rsspsrv "github.com/lob-inc/rssp/server"
)

func main() {
	flag.Usage = func() {
		fmt.Fprint(os.Stderr, "Usage:\n")
		fmt.Fprint(os.Stderr, "  api [OPTIONS]\n\n")
		fmt.Fprintln(os.Stderr, flag.CommandLine.FlagUsages())
	}
	flag.Parse()

	// if srv.Port == 0 {
	// 	srv.Port = conf.Server.APIPort
	// }
	rsspConf := rsspsrv.NewConfig()
	srv, err := newServer(rsspConf)
	if err != nil {
		panic(err)
	}
	srv.srv.ListenAndServe()
}
