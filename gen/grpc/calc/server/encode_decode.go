// Code generated by goa v3.0.10, DO NOT EDIT.
//
// calc gRPC server encoders and decoders
//
// Command:
// $ goa gen calc/design

package server

import (
	calc "calc/gen/calc"
	calcpb "calc/gen/grpc/calc/pb"
	"context"

	goagrpc "goa.design/goa/v3/grpc"
	"google.golang.org/grpc/metadata"
)

// EncodeAddResponse encodes responses from the "calc" service "add" endpoint.
func EncodeAddResponse(ctx context.Context, v interface{}, hdr, trlr *metadata.MD) (interface{}, error) {
	result, ok := v.(int)
	if !ok {
		return nil, goagrpc.ErrInvalidType("calc", "add", "int", v)
	}
	resp := NewAddResponse(result)
	return resp, nil
}

// DecodeAddRequest decodes requests sent to "calc" service "add" endpoint.
func DecodeAddRequest(ctx context.Context, v interface{}, md metadata.MD) (interface{}, error) {
	var (
		message *calcpb.AddRequest
		ok      bool
	)
	{
		if message, ok = v.(*calcpb.AddRequest); !ok {
			return nil, goagrpc.ErrInvalidType("calc", "add", "*calcpb.AddRequest", v)
		}
	}
	var payload *calc.AddPayload
	{
		payload = NewAddPayload(message)
	}
	return payload, nil
}

// EncodeRestaResponse encodes responses from the "calc" service "resta"
// endpoint.
func EncodeRestaResponse(ctx context.Context, v interface{}, hdr, trlr *metadata.MD) (interface{}, error) {
	result, ok := v.(int)
	if !ok {
		return nil, goagrpc.ErrInvalidType("calc", "resta", "int", v)
	}
	resp := NewRestaResponse(result)
	return resp, nil
}

// DecodeRestaRequest decodes requests sent to "calc" service "resta" endpoint.
func DecodeRestaRequest(ctx context.Context, v interface{}, md metadata.MD) (interface{}, error) {
	var (
		message *calcpb.RestaRequest
		ok      bool
	)
	{
		if message, ok = v.(*calcpb.RestaRequest); !ok {
			return nil, goagrpc.ErrInvalidType("calc", "resta", "*calcpb.RestaRequest", v)
		}
	}
	var payload *calc.RestaPayload
	{
		payload = NewRestaPayload(message)
	}
	return payload, nil
}

// EncodeMultiplicacionResponse encodes responses from the "calc" service
// "multiplicacion" endpoint.
func EncodeMultiplicacionResponse(ctx context.Context, v interface{}, hdr, trlr *metadata.MD) (interface{}, error) {
	result, ok := v.(int)
	if !ok {
		return nil, goagrpc.ErrInvalidType("calc", "multiplicacion", "int", v)
	}
	resp := NewMultiplicacionResponse(result)
	return resp, nil
}

// DecodeMultiplicacionRequest decodes requests sent to "calc" service
// "multiplicacion" endpoint.
func DecodeMultiplicacionRequest(ctx context.Context, v interface{}, md metadata.MD) (interface{}, error) {
	var (
		message *calcpb.MultiplicacionRequest
		ok      bool
	)
	{
		if message, ok = v.(*calcpb.MultiplicacionRequest); !ok {
			return nil, goagrpc.ErrInvalidType("calc", "multiplicacion", "*calcpb.MultiplicacionRequest", v)
		}
	}
	var payload *calc.MultiplicacionPayload
	{
		payload = NewMultiplicacionPayload(message)
	}
	return payload, nil
}

// EncodeDivisionResponse encodes responses from the "calc" service "division"
// endpoint.
func EncodeDivisionResponse(ctx context.Context, v interface{}, hdr, trlr *metadata.MD) (interface{}, error) {
	result, ok := v.(int)
	if !ok {
		return nil, goagrpc.ErrInvalidType("calc", "division", "int", v)
	}
	resp := NewDivisionResponse(result)
	return resp, nil
}

// DecodeDivisionRequest decodes requests sent to "calc" service "division"
// endpoint.
func DecodeDivisionRequest(ctx context.Context, v interface{}, md metadata.MD) (interface{}, error) {
	var (
		message *calcpb.DivisionRequest
		ok      bool
	)
	{
		if message, ok = v.(*calcpb.DivisionRequest); !ok {
			return nil, goagrpc.ErrInvalidType("calc", "division", "*calcpb.DivisionRequest", v)
		}
	}
	var payload *calc.DivisionPayload
	{
		payload = NewDivisionPayload(message)
	}
	return payload, nil
}
