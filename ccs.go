package gnarkserializable

import (
	"io"

	"github.com/consensys/gnark/backend/groth16"
	"github.com/consensys/gnark/constraint"
)

type ConstraintSystem struct{ constraint.ConstraintSystem }

func (ccs *ConstraintSystem) IsNil() bool {
	return ccs == nil
}

func (ccs *ConstraintSystem) WriteTo(writer io.Writer) (int64, error) {
	return ccs.ConstraintSystem.WriteTo(writer)
}

func (ccs *ConstraintSystem) ReadFrom(reader io.Reader) (int64, error) {
	return ccs.ConstraintSystem.ReadFrom(reader)
}

func (ccs *ConstraintSystem) MarshalJSON() ([]byte, error) {
	return WriteTo(ccs)
}

func (ccs *ConstraintSystem) UnmarshalJSON(buf []byte) error {
	if ccs == nil {
		ccs = &ConstraintSystem{}
	}

	if ccs.ConstraintSystem == nil {
		ccs.ConstraintSystem = groth16.NewCS(EcCurve)
	}

	return ReadFrom(buf, ccs)
}
