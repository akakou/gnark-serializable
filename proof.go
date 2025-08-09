package gnarkserializable

import (
	"io"

	"github.com/consensys/gnark/backend/groth16"
)

type Proof struct{ Proof groth16.Proof }

func (proof *Proof) IsNil() bool {
	return proof == nil
}

func (proof *Proof) WriteTo(writer io.Writer) (int64, error) {
	return proof.Proof.WriteTo(writer)
}

func (proof *Proof) ReadFrom(reader io.Reader) (int64, error) {
	return proof.Proof.ReadFrom(reader)
}

func (proof *Proof) MarshalJSON() ([]byte, error) {
	return WriteTo(proof)
}

func (proof *Proof) UnmarshalJSON(buf []byte) error {
	if proof == nil {
		proof = &Proof{}
	}

	if proof.Proof == nil {
		proof.Proof = groth16.NewProof(EcCurve)
	}

	return ReadFrom(buf, proof)
}
