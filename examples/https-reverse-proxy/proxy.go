package main

import (
	"github.com/yeqown/log"

	"github.com/valyala/fasthttp"
	proxy "github.com/yeqown/fasthttp-reverse-proxy"
)

var (
	proxyServer = proxy.NewReverseProxy("localhost:8080", proxy.WithTLS())
)

// ProxyHandler ... fasthttp.RequestHandler func
func ProxyHandler(ctx *fasthttp.RequestCtx) {
	requestURI := string(ctx.RequestURI())
	log.Info("a request incoming, requestURI=", requestURI)
	proxyServer.ServeHTTP(ctx)
}

func main() {
	if err := fasthttp.ListenAndServe("0.0.0.0:8081", ProxyHandler); err != nil {
		log.Fatal(err)
	}
}
