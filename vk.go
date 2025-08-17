package gnarkserializable

import (
	"io"

	"github.com/consensys/gnark/backend/groth16"
)

type VerifyingKey struct{ VerifyingKey groth16.VerifyingKey }

func (vk *VerifyingKey) IsNil() bool {
	return vk == nil
}

func (vk VerifyingKey) WriteTo(writer io.Writer) (int64, error) {
	return Write(vk.VerifyingKey, writer)
}

func (vk *VerifyingKey) ReadFrom(reader io.Reader) (int64, error) {
	return Read(vk.VerifyingKey, reader)
}

func (vk *VerifyingKey) MarshalJSON() ([]byte, error) {
	return WriteTo(vk)
}

func (vk *VerifyingKey) UnmarshalJSON(buf []byte) error {
	if vk == nil {
		vk = &VerifyingKey{}
	}

	if vk.VerifyingKey == nil {
		vk.VerifyingKey = groth16.NewVerifyingKey(EcCurve)
	}

	return ReadFrom(buf, vk)
}
