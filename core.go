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

func WriteTo[T Writable](data T) ([]byte, error) {
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

func ReadFrom[T Readable](buf []byte, data T) error {
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
