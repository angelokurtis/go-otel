//go:generate go-enum --nocase --names

package env

// Compression defines the compression type to use on OTLP.
// ENUM(none, gzip)
type Compression string

func (x *Compression) UnmarshalText(text []byte) error {
	var err error
	*x, err = ParseCompression(string(text))

	return err
}

// Protocol defines the encoding of telemetry data and the protocol used to exchange data between the client and the server.
// ENUM(grpc, http/protobuf)
type Protocol string

func (x *Protocol) UnmarshalText(text []byte) error {
	var err error
	*x, err = ParseProtocol(string(text))

	return err
}
