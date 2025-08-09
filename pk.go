package gnarkserializable

import (
	"io"

	"github.com/consensys/gnark/backend/groth16"
)

type ProvingKey struct{ ProvingKey groth16.ProvingKey }

func (pk ProvingKey) IsNil() bool {
	return pk.ProvingKey == nil
}

func (pk ProvingKey) WriteTo(writer io.Writer) (int64, error) {
	return pk.ProvingKey.WriteTo(writer)
}

func (pk *ProvingKey) ReadFrom(reader io.Reader) (int64, error) {
	return pk.ProvingKey.ReadFrom(reader)
}

func (pk ProvingKey) MarshalJSON() ([]byte, error) {
	return WriteTo(pk)
}

func (pk *ProvingKey) UnmarshalJSON(buf []byte) error {
	if pk == nil {
		pk = &ProvingKey{}
	}

	if pk.ProvingKey == nil {
		pk.ProvingKey = groth16.NewProvingKey(EcCurve)
	}

	return ReadFrom(buf, pk)
}
