package rpc

import (
	"google.golang.org/grpc/codes"

	"github.com/negz/practice/tiny/proto"
)

// Error handling patterns inspired by the following article:
// https://dave.cheney.net/2016/04/27/dont-just-check-errors-handle-them-gracefully

type codedError struct {
	error
	code codes.Code
}

func (err *codedError) Code() codes.Code {
	return err.code
}

// Codify associates the supplied gRPC error code with the given error.
// The returned error will produce the supplied gRPC error code when processed
// by ToStatus. Codify returns nil when passed a nil error.
// Note this is almost https://godoc.org/google.golang.org/grpc#Errorf, except
// that it embeds the original error and thus preserves it.
func Codify(err error, code codes.Code) error {
	if err == nil {
		return nil
	}
	return &codedError{err, code}
}

// Status converts the given error into a status response.
// If the error is nil a status with codes.OK is returned.
// If the error has a code attached it will be returned.
// Failing the above a status with codes.Unknown will be returned.
func Status(err error) *proto.Status {
	// No error. Return an OK status (codes.OK is the uint32 zero value).
	if err == nil {
		return &proto.Status{}
	}

	type coder interface {
		Code() codes.Code
	}

	// Error has a code attached.
	if e, ok := err.(coder); ok {
		return &proto.Status{Code: uint32(e.Code()), Message: err.Error()}
	}

	// Unable to automatically determine error.
	return &proto.Status{Code: uint32(codes.Unknown), Message: err.Error()}
}

// IsNotFound determines whether an error indicates something was not found.
// It does this by walking down the stack of errors built by pkg/errors and
// returning true for the first error that implements the following interface:
//
// type notfounder interface {
//   NotFound()
// }
func IsNotFound(err error) bool {
	if err == nil {
		return false
	}
	for {
		if _, ok := err.(interface {
			NotFound()
		}); ok {
			return true
		}
		if c, ok := err.(interface {
			Cause() error
		}); ok {
			err = c.Cause()
			continue
		}
		return false
	}
}
