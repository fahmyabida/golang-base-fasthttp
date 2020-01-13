package routing

import (
	"flag"
	"fmt"
	"golang-labbaika_gaji-fasthttp/model"
	"golang-labbaika_gaji-fasthttp/util"
	"time"

	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
)

//GatewayHandler is struct for handler gateway router in fastHTTP
type GatewayHandler struct {
	properties model.Properties
}

//InitRouting is used to init routing on nats that will handle a request from client
func InitRouting(properties model.Properties, timeOut time.Duration) GatewayHandler {
	return GatewayHandler{properties}
}

//Routing is contain routing path
func (h *GatewayHandler) Routing() {
	var (
		addr = flag.String("addr", ":"+h.properties.PortServer, "TCP address to listen to")
	)

	router := fasthttprouter.New()

	router.GET("/hello", h.helloWorld)

	util.Event("Listening on " + *addr)
	fmt.Println("Listening on " + *addr)
	fmt.Println("Ready to serve ~")
	if errorIntialRunning := fasthttp.ListenAndServe(*addr, CORS(router.Handler)); errorIntialRunning != nil {
		util.Error("Error in ListenAndServe " + errorIntialRunning.Error())
		fmt.Println("Failed to serve ~")
	}
}

func (h *GatewayHandler) helloWorld(ctx *fasthttp.RequestCtx) {
	fmt.Fprintln(ctx, "hello its me")
}

