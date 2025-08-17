package gnarkserializable_test

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"

	gnarkserializable "github.com/akakou/gnark-serializable"
	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark/backend/groth16"
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/frontend/cs/r1cs"
)

type Circuit struct {
}

func (circuit *Circuit) Define(api frontend.API) error {
	return nil
}

type Json struct {
	Proof gnarkserializable.Proof
	CCS   gnarkserializable.ConstraintSystem
	PK    gnarkserializable.ProvingKey
	VK    gnarkserializable.VerifyingKey
	A     int
}

func TestAll(t *testing.T) {
	gnarkserializable.Unsafe = true
	var circuit Circuit
	r1cs, err := frontend.Compile(
		ecc.BLS12_381.ScalarField(),
		r1cs.NewBuilder,
		&circuit)
	if err != nil {
		log.Fatalf("compile: %v", err)
	}

	gpk, gvk, err := groth16.Setup(r1cs)
	if err != nil {
		log.Fatalf("setup: %v", err)
	}

	assignment := &Circuit{}
	witness, err := frontend.NewWitness(assignment, ecc.BLS12_381.ScalarField())
	if err != nil {
		log.Fatalf("witness: %v", err)
	}

	proof, err := groth16.Prove(r1cs, gpk, witness)
	if err != nil {
		log.Fatalf("prove: %v", err)
	}

	var res1 Json
	proof2 := gnarkserializable.Proof{proof}
	ccs := gnarkserializable.ConstraintSystem{r1cs}
	pk := gnarkserializable.ProvingKey{gpk}
	vk := gnarkserializable.VerifyingKey{gvk}

	buf, err := json.Marshal(&Json{
		Proof: proof2,
		A:     100,
		CCS:   ccs,
		PK:    pk,
		VK:    vk,
	})
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(buf, &res1)
	if err != nil {
		panic(err)
	}

	res2, err := json.Marshal(&res1)
	if err != nil {
		panic(err)
	}

	if string(buf) != string(res2) {
		fmt.Printf("\n-----\nresult: %v\n-----\n%v\n", string(buf), string(res2))
		panic("not match")
	}
}
