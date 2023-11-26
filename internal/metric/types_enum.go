// Code generated by go-enum DO NOT EDIT.
// Version: 0.5.6
// Revision: 97611fddaa414f53713597918c5e954646cb8623
// Build Date: 2023-03-26T21:38:06Z
// Built By: goreleaser

package metric

import (
	"errors"
	"fmt"
	"strings"
)

const (
	// CompressionNone is a Compression of type none.
	CompressionNone Compression = "none"
	// CompressionGzip is a Compression of type gzip.
	CompressionGzip Compression = "gzip"
)

var ErrInvalidCompression = errors.New("not a valid Compression")

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
	// ExporterOtlp is a Exporter of type otlp.
	ExporterOtlp Exporter = "otlp"
	// ExporterNone is a Exporter of type none.
	ExporterNone Exporter = "none"
	// ExporterPrometheus is a Exporter of type prometheus.
	ExporterPrometheus Exporter = "prometheus"
	// ExporterLogging is a Exporter of type logging.
	ExporterLogging Exporter = "logging"
)

var ErrInvalidExporter = errors.New("not a valid Exporter")

// String implements the Stringer interface.
func (x Exporter) String() string {
	return string(x)
}

// IsValid provides a quick way to determine if the typed value is
// part of the allowed enumerated values
func (x Exporter) IsValid() bool {
	_, err := ParseExporter(string(x))
	return err == nil
}

var _ExporterValue = map[string]Exporter{
	"otlp":       ExporterOtlp,
	"none":       ExporterNone,
	"prometheus": ExporterPrometheus,
	"logging":    ExporterLogging,
}

// ParseExporter attempts to convert a string to a Exporter.
func ParseExporter(name string) (Exporter, error) {
	if x, ok := _ExporterValue[name]; ok {
		return x, nil
	}
	// Case insensitive parse, do a separate lookup to prevent unnecessary cost of lowercasing a string if we don't need to.
	if x, ok := _ExporterValue[strings.ToLower(name)]; ok {
		return x, nil
	}

	return Exporter(""), fmt.Errorf("%s is %w", name, ErrInvalidExporter)
}

const (
	// ProtocolGrpc is a Protocol of type grpc.
	ProtocolGrpc Protocol = "grpc"
	// ProtocolHttpProtobuf is a Protocol of type http/protobuf.
	ProtocolHttpProtobuf Protocol = "http/protobuf"
)

var ErrInvalidProtocol = errors.New("not a valid Protocol")

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
