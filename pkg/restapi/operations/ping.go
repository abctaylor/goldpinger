// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// PingHandlerFunc turns a function with the right signature into a ping handler
type PingHandlerFunc func(PingParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PingHandlerFunc) Handle(params PingParams) middleware.Responder {
	return fn(params)
}

// PingHandler interface for that can handle valid ping params
type PingHandler interface {
	Handle(PingParams) middleware.Responder
}

// NewPing creates a new http.Handler for the ping operation
func NewPing(ctx *middleware.Context, handler PingHandler) *Ping {
	return &Ping{Context: ctx, Handler: handler}
}

/*Ping swagger:route GET /ping ping

return query stats

*/
type Ping struct {
	Context *middleware.Context
	Handler PingHandler
}

func (o *Ping) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPingParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
