package main

import (
	"flag"
	"github.com/facebookgo/grace/gracehttp"
	"net/http"
	"softnet-agent/handler"
)

var (
	address0 = flag.String("a0", ":9999","")
)

func main()  {
	flag.Parse()
	gracehttp.Serve(
		&http.Server{Addr: *address0, Handler: handler.HandlerInit()},
	)
}
