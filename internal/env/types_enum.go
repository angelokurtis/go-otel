// Code generated by go-enum DO NOT EDIT.
// Version: 0.5.6
// Revision: 97611fddaa414f53713597918c5e954646cb8623
// Build Date: 2023-03-26T21:38:06Z
// Built By: goreleaser

package env

import (
	"fmt"
	"strings"
)

const (
	// CompressionNone is a Compression of type none.
	CompressionNone Compression = "none"
	// CompressionGzip is a Compression of type gzip.
	CompressionGzip Compression = "gzip"
)

var ErrInvalidCompression = fmt.Errorf("not a valid Compression, try [%s]", strings.Join(_CompressionNames, ", "))

var _CompressionNames = []string{
	string(CompressionNone),
	string(CompressionGzip),
}

// CompressionNames returns a list of possible string values of Compression.
func CompressionNames() []string {
	tmp := make([]string, len(_CompressionNames))
	copy(tmp, _CompressionNames)
	return tmp
}

// String implements the Stringer interface.
func (x Compression) String() string {
	return string(x)
}

// IsValid provides a quick way to determine if the typed value is
// part of the allowed enumerated values
func (x Compression) IsValid() bool {
	_, err := ParseCompression(string(x))
	return err == nil
}

var _CompressionValue = map[string]Compression{
	"none": CompressionNone,
	"gzip": CompressionGzip,
}

// ParseCompression attempts to convert a string to a Compression.
func ParseCompression(name string) (Compression, error) {
	if x, ok := _CompressionValue[name]; ok {
		return x, nil
	}
	// Case insensitive parse, do a separate lookup to prevent unnecessary cost of lowercasing a string if we don't need to.
	if x, ok := _CompressionValue[strings.ToLower(name)]; ok {
		return x, nil
	}
	return Compression(""), fmt.Errorf("%s is %w", name, ErrInvalidCompression)
}

const (
	// ProtocolGrpc is a Protocol of type grpc.
	ProtocolGrpc Protocol = "grpc"
	// ProtocolHttpProtobuf is a Protocol of type http/protobuf.
	ProtocolHttpProtobuf Protocol = "http/protobuf"
)

var ErrInvalidProtocol = fmt.Errorf("not a valid Protocol, try [%s]", strings.Join(_ProtocolNames, ", "))

var _ProtocolNames = []string{
	string(ProtocolGrpc),
	string(ProtocolHttpProtobuf),
}

// ProtocolNames returns a list of possible string values of Protocol.
func ProtocolNames() []string {
	tmp := make([]string, len(_ProtocolNames))
	copy(tmp, _ProtocolNames)
	return tmp
}

// String implements the Stringer interface.
func (x Protocol) String() string {
	return string(x)
}

// IsValid provides a quick way to determine if the typed value is
// part of the allowed enumerated values
func (x Protocol) IsValid() bool {
	_, err := ParseProtocol(string(x))
	return err == nil
}

var _ProtocolValue = map[string]Protocol{
	"grpc":          ProtocolGrpc,
	"http/protobuf": ProtocolHttpProtobuf,
}

// ParseProtocol attempts to convert a string to a Protocol.
func ParseProtocol(name string) (Protocol, error) {
	if x, ok := _ProtocolValue[name]; ok {
		return x, nil
	}
	// Case insensitive parse, do a separate lookup to prevent unnecessary cost of lowercasing a string if we don't need to.
	if x, ok := _ProtocolValue[strings.ToLower(name)]; ok {
		return x, nil
	}
	return Protocol(""), fmt.Errorf("%s is %w", name, ErrInvalidProtocol)
}