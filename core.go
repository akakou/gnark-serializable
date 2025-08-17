package gnarkserializable

import (
	"bytes"
	"encoding/base64"
	"io"
	"strconv"

	"github.com/consensys/gnark-crypto/ecc"
)

var EcCurve = ecc.BLS12_381

type Writable interface {
	WriteTo(io.Writer) (int64, error)
	IsNil() bool
}

type Readable interface {
	ReadFrom(io.Reader) (int64, error)
}

func WriteTo(data Writable) ([]byte, error) {
	if data.IsNil() {
		return []byte("null"), nil
	}
	var buffer bytes.Buffer
	_, err := data.WriteTo(&buffer)
	if err != nil {
		return nil, err
	}

	raw := buffer.Bytes()
	enc := base64.URLEncoding.EncodeToString(raw)
	res := strconv.Quote(enc)

	return []byte(res), nil
}

func ReadFrom(buf []byte, data Readable) error {
	if string(buf) == "null" {
		return nil
	}

	unq, err := strconv.Unquote(string(buf))
	if err != nil {
		return err
	}

	raw, err := base64.URLEncoding.DecodeString(unq)
	if err != nil {
		return err
	}

	reader := bytes.NewReader(raw)
	_, err = data.ReadFrom(reader)

	return err
}

type UnsafeWritable interface {
	WriteTo(io.Writer) (int64, error)
	WriteRawTo(io.Writer) (int64, error)
}

type UnsafeReadable interface {
	ReadFrom(io.Reader) (int64, error)
	UnsafeReadFrom(io.Reader) (int64, error)
}

var Unsafe = false

func Write(unsafe UnsafeWritable, io io.Writer) (int64, error) {
	if Unsafe {
		return unsafe.WriteRawTo(io)
	} else {
		return unsafe.WriteTo(io)
	}
}

func Read(unsafe UnsafeReadable, io io.Reader) (int64, error) {
	if Unsafe {
		return unsafe.UnsafeReadFrom(io)
	} else {
		return unsafe.ReadFrom(io)
	}
}
