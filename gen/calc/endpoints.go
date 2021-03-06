// Code generated by goa v3.0.10, DO NOT EDIT.
//
// calc endpoints
//
// Command:
// $ goa gen calc/design

package calc

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Endpoints wraps the "calc" service endpoints.
type Endpoints struct {
	Add            goa.Endpoint
	Resta          goa.Endpoint
	Multiplicacion goa.Endpoint
	Division       goa.Endpoint
}

// NewEndpoints wraps the methods of the "calc" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	return &Endpoints{
		Add:            NewAddEndpoint(s),
		Resta:          NewRestaEndpoint(s),
		Multiplicacion: NewMultiplicacionEndpoint(s),
		Division:       NewDivisionEndpoint(s),
	}
}

// Use applies the given middleware to all the "calc" service endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.Add = m(e.Add)
	e.Resta = m(e.Resta)
	e.Multiplicacion = m(e.Multiplicacion)
	e.Division = m(e.Division)
}

// NewAddEndpoint returns an endpoint function that calls the method "add" of
// service "calc".
func NewAddEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*AddPayload)
		return s.Add(ctx, p)
	}
}

// NewRestaEndpoint returns an endpoint function that calls the method "resta"
// of service "calc".
func NewRestaEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*RestaPayload)
		return s.Resta(ctx, p)
	}
}

// NewMultiplicacionEndpoint returns an endpoint function that calls the method
// "multiplicacion" of service "calc".
func NewMultiplicacionEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*MultiplicacionPayload)
		return s.Multiplicacion(ctx, p)
	}
}

// NewDivisionEndpoint returns an endpoint function that calls the method
// "division" of service "calc".
func NewDivisionEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*DivisionPayload)
		return s.Division(ctx, p)
	}
}
