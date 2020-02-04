package calcapi

import (
	calc "calc/gen/calc"
	"context"
	"log"
)

// calc service example implementation.
// The example methods log the requests and return zero values.
type calcsrvc struct {
	logger *log.Logger
}

// NewCalc returns the calc service implementation.
func NewCalc(logger *log.Logger) calc.Service {
	return &calcsrvc{logger}
}

// Add implements add.
func (s *calcsrvc) Add(ctx context.Context, p *calc.AddPayload) (res int, err error) {
	s.logger.Print("calc.add")
	return  p.A + p.B , nil
}

// Resta implements resta.
func (s *calcsrvc) Resta(ctx context.Context, p *calc.RestaPayload) (res int, err error) {
	s.logger.Print("calc.resta")
	return p.A - p.B, nil
}

// Multiplicacion implements multiplicacion.
func (s *calcsrvc) Multiplicacion(ctx context.Context, p *calc.MultiplicacionPayload) (res int, err error) {
	s.logger.Print("calc.multiplicacion")
	return p.A * p.B, nil
}

// Division implements division.
func (s *calcsrvc) Division(ctx context.Context, p *calc.DivisionPayload) (res int, err error) {
	s.logger.Print("calc.division")
	return p.A / p.B , nil
}
