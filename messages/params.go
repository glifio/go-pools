package messages

import (
	"bytes"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	cborutil "github.com/filecoin-project/go-cbor-util"
	"github.com/filecoin-project/go-state-types/builtin"
	"github.com/filecoin-project/lotus/chain/types"
	"github.com/glifio/go-pools/abigen"
)

type MethodLookupError struct {
	Err error
}

func (r *MethodLookupError) Error() string {
	return fmt.Sprintf("%v", r.Err)
}

func ParseAgentParams(msg types.MessageTrace) (*abi.Method, map[string]interface{}, error) {
	if msg.Method != builtin.MustGenerateFRCMethodNum("InvokeEVM") {
		return nil, nil, nil
	}
	fmt.Printf("Jim0 msg %+v\n", msg)
	data := msg.Params

	if len(data) == 0 {
		return nil, nil, nil
	}

	var paramsBytes []byte = data

	sig := paramsBytes[0:4]

	abi, err := abigen.AgentMetaData.GetAbi()
	if err != nil {
		return nil, nil, err
	}

	fmt.Println("Jim1")
	method, err := abi.MethodById(sig)
	if err != nil {
		// Try unpacking the params as CBOR
		reader := bytes.NewReader(data)
		err := cborutil.ReadCborRPC(reader, &paramsBytes)
		if err != nil {
			return nil, nil, &MethodLookupError{Err: err}
		}
		sig = paramsBytes[0:4]
		method, err = abi.MethodById(sig)
		if err != nil {
			return nil, nil, &MethodLookupError{Err: err}
		}
	}
	fmt.Println("Jim2")

	unpackedMap := make(map[string]interface{})
	err = method.Inputs.UnpackIntoMap(unpackedMap, paramsBytes[4:])
	if err != nil {
		return nil, nil, err
	}
	fmt.Println("Jim3")

	return method, unpackedMap, nil
}

func GetSignedCredentialFromParams(params map[string]interface{}) (*abigen.SignedCredential, error) {
	var signedCred *abigen.SignedCredential
	scParams, ok := params["sc"]
	if !ok {
		return nil, nil
	}
	sc, _ := scParams.(struct {
		Vc struct {
			Issuer          common.Address "json:\"issuer\""
			Subject         *big.Int       "json:\"subject\""
			EpochIssued     *big.Int       "json:\"epochIssued\""
			EpochValidUntil *big.Int       "json:\"epochValidUntil\""
			Value           *big.Int       "json:\"value\""
			Action          [4]uint8       "json:\"action\""
			Target          uint64         "json:\"target\""
			Claim           []uint8        "json:\"claim\""
		} "json:\"vc\""
		V uint8     "json:\"v\""
		R [32]uint8 "json:\"r\""
		S [32]uint8 "json:\"s\""
	})
	signedCred = &abigen.SignedCredential{
		Vc: abigen.VerifiableCredential{
			Issuer:          sc.Vc.Issuer,
			Subject:         sc.Vc.Subject,
			EpochIssued:     sc.Vc.EpochIssued,
			EpochValidUntil: sc.Vc.EpochValidUntil,
			Value:           sc.Vc.Value,
			Action:          sc.Vc.Action,
			Target:          sc.Vc.Target,
			Claim:           sc.Vc.Claim,
		},
		V: sc.V,
		R: sc.R,
		S: sc.S,
	}
	return signedCred, nil
}
