package main

import (
	"context"
	"log"
	"net/http"
	"strings"

	loads "github.com/go-openapi/loads"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	rsspsrv "github.com/lob-inc/rssp/server"
	"github.com/lob-inc/rssp/server/api/restapi"
	"github.com/lob-inc/rssp/server/api/restapi/operations"
	"github.rakops.com/gatd/rad.api.pb/gen/go/adsapipb"
	"google.golang.org/grpc"
)

type server struct {
	srv        *http.Server
	rsspServer *restapi.Server

	mux *http.ServeMux

	ctx    context.Context
	cancel context.CancelFunc
}

func newServer(rsspConf *rsspsrv.Config) (*server, error) {
	rmux := runtime.NewServeMux()
	ctx, cancel := context.WithCancel(context.Background())

	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := adsapipb.RegisterAdsHandlerFromEndpoint(ctx, rmux, "localhost:6012", opts)
	if err != nil {
		return nil, err
	}

	mux := http.NewServeMux()
	rsspHandler := newRSSPServer(rsspConf)
	mux.Handle("/", rsspHandler)
	mux.Handle("/ads", http.StripPrefix("/ads", rmux))

	srv := &http.Server{
		Addr: ":8080",
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			p := r.URL.Path
			switch {
			case strings.HasPrefix(p, "/v1/ads"):
				r.URL.Path = strings.Replace(r.URL.Path, "/v1/ads", "", 1)
				rmux.ServeHTTP(w, r)
			case strings.HasPrefix(p, "/v1"):
				rsspHandler.ServeHTTP(w, r)
			}
		}),
	}
	return &server{
		srv:    srv,
		ctx:    ctx,
		cancel: cancel,
	}, nil
}

func newRSSPServer(conf *rsspsrv.Config) http.Handler {
	swaggerSpec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(adsapipb.GetCampaignsRequest{})

	var srv *restapi.Server // make sure init is called

	api := operations.NewRsspAPI(swaggerSpec)
	srv = restapi.NewServer(api)

	restapi.InjectServiceDependencies(api, conf)
	operations.ConfigureServices(api)
	srv.ConfigureAPI()
	return srv.GetHandler()
}
