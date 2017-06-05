package protorec

import (
	"io"

	"github.com/dedis/protobuf"
	"github.com/uluyol/binrec"
)

// This file contains helper io functions compatible with
// Java protobuf's writeDelimitedTo and mergeDelimitedFrom methods.

// WriteDelimitedTo writes a message to an io.Writer in a varint-delimited format.
//
// WriteDelimitedTo is analogous to writeDelimitedTo in protobuf-java.
func WriteDelimitedTo(w io.Writer, m interface{}) error {
	data, err := protobuf.Encode(m)

	if err != nil {
		return err
	}
	return binrec.WriteDelimitedTo(w, data)
}

type Reader interface {
	io.ByteReader
	io.Reader
}

// ReadDelimitedFrom reads a message from a Reader in a varint-delimited format.
// bufio.Reader may be used to construct a Reader.
//
// ReadDelimitedFrom is analogous to mergeDelimitedFrom in protobuf-java.
func ReadDelimitedFrom(r Reader, m interface{}) error {
	buf, err := binrec.ReadDelimitedFrom(r)
	if err != nil {
		return err
	}
	return protobuf.Decode(buf, m)
}
