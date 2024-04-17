package vc

import (
	"errors"
	"math/big"
	"reflect"

	"github.com/glifio/go-pools/abigen"
	"github.com/jimpick/go-ethereum/accounts/abi"
	"github.com/jimpick/go-ethereum/common"
	"github.com/jimpick/go-ethereum/common/hexutil"
	"github.com/jimpick/go-ethereum/common/math"
	"github.com/jimpick/go-ethereum/crypto"
	"github.com/jimpick/go-ethereum/signer/core/apitypes"
)

var (
	ErrZeroSubject = errors.New("subject cannot be 0")
)

type AgentData struct {
	AgentValue                  *big.Int
	CollateralValue             *big.Int
	ExpectedDailyFaultPenalties *big.Int
	ExpectedDailyRewards        *big.Int
	Gcred                       *big.Int
	QaPower                     *big.Int
	Principal                   *big.Int
	FaultySectors               *big.Int
	LiveSectors                 *big.Int
	GreenScore                  *big.Int
}

func TypedEIP712(
	chainId int64,
	verifyingContract common.Address,
	vc *abigen.VerifiableCredential,
) apitypes.TypedData {
	return apitypes.TypedData{
		Types: apitypes.Types{
			"VerifiableCredential": []apitypes.Type{
				// the 0x address of the vc issuer
				{Name: "issuer", Type: "address"},
				// the ID of the agent requesting the credential
				{Name: "subject", Type: "uint256"},
				// the current epoch height
				{Name: "epochIssued", Type: "uint256"},
				// the epoch this vc is valid until
				{Name: "epochValidUntil", Type: "uint256"},
				// any value corresponding to this VC request
				{Name: "value", Type: "uint256"},
				// the function signature of the method to call, passing this VC as an argument
				{Name: "action", Type: "bytes4"},
				// the ID address of the built-in miner actor corresponding to the action of this VC
				{Name: "target", Type: "uint64"},
				// encoded claim data
				{Name: "claim", Type: "bytes"},
			},
			"EIP712Domain": []apitypes.Type{
				{Name: "name", Type: "string"},
				{Name: "version", Type: "string"},
				{Name: "chainId", Type: "uint256"},
				{Name: "verifyingContract", Type: "address"},
			},
		},
		PrimaryType: "VerifiableCredential",
		Domain: apitypes.TypedDataDomain{
			Name:              "glif.io",
			Version:           "1",
			ChainId:           math.NewHexOrDecimal256(chainId),
			VerifyingContract: verifyingContract.String(),
		},
		Message: apitypes.TypedDataMessage{
			"issuer":          vc.Issuer.String(),
			"subject":         vc.Subject.String(),
			"epochIssued":     vc.EpochIssued.String(),
			"epochValidUntil": vc.EpochValidUntil.String(),
			"value":           vc.Value.String(),
			"action":          vc.Action,
			"target":          hexutil.EncodeUint64(vc.Target),
			"claim":           vc.Claim,
		},
	}
}

func AbiEncodeClaim(claim AgentData) ([]byte, error) {
	// here we create an array of Args to encode for AgentData
	var arguments abi.Arguments
	// here are the params to abi.encode into bytes
	var params []interface{}
	// the uint256 abi type to encode the AgentData with
	uint256Ty, _ := abi.NewType("uint256", "", nil)

	values := reflect.ValueOf(claim)
	types := values.Type()

	for i := 0; i < values.NumField(); i++ {
		fieldType := uint256Ty
		// add the uint256 type and the key name to the AgentData abi
		arguments = append(arguments, abi.Argument{
			Type: fieldType,
			Name: types.Field(i).Name,
		})
		params = append(params, values.Field(i).Interface())
	}
	// abi.encode the AgentData
	return arguments.Pack(params...)
}

func AbiDecodeClaim(claim []byte) (*AgentData, error) {
	// Define the ABI arguments
	var arguments abi.Arguments
	uint256Ty, _ := abi.NewType("uint256", "", nil)
	numFields := reflect.TypeOf(AgentData{}).NumField()

	for i := 0; i < numFields; i++ {
		arguments = append(arguments, abi.Argument{Type: uint256Ty})
	}

	out, err := arguments.Unpack(claim)
	if err != nil {
		return nil, err
	}

	// Convert the slice of big.Ints into an AgentData
	return &AgentData{
		AgentValue:                  out[0].(*big.Int),
		CollateralValue:             out[1].(*big.Int),
		ExpectedDailyFaultPenalties: out[2].(*big.Int),
		ExpectedDailyRewards:        out[3].(*big.Int),
		Gcred:                       out[4].(*big.Int),
		QaPower:                     out[5].(*big.Int),
		Principal:                   out[6].(*big.Int),
		FaultySectors:               out[7].(*big.Int),
		LiveSectors:                 out[8].(*big.Int),
		GreenScore:                  out[9].(*big.Int),
	}, nil
}

func NewVerifiableCredential(
	issuer common.Address,
	subject *big.Int,
	epochIssued *big.Int,
	epochValidUntil *big.Int,
	value *big.Int,
	action [4]byte,
	target uint64,
	claim AgentData,
) (*abigen.VerifiableCredential, error) {
	// no 0 subjects allowed (this is protect against on the contracts, but easier to throw a more helpful error here)
	if subject.Cmp(common.Big0) == 0 {
		return &abigen.VerifiableCredential{}, ErrZeroSubject
	}
	claimBytes, err := AbiEncodeClaim(claim)
	if err != nil {
		return &abigen.VerifiableCredential{}, err
	}

	return &abigen.VerifiableCredential{
		Issuer:          issuer,
		Subject:         subject,
		EpochIssued:     epochIssued,
		EpochValidUntil: epochValidUntil,
		Value:           value,
		Action:          action,
		Target:          target,
		Claim:           claimBytes,
	}, nil
}

func EmptyAgentData() *AgentData {
	return &AgentData{
		AgentValue:                  big.NewInt(0),
		CollateralValue:             big.NewInt(0),
		ExpectedDailyFaultPenalties: big.NewInt(0),
		ExpectedDailyRewards:        big.NewInt(0),
		Gcred:                       big.NewInt(100),
		QaPower:                     big.NewInt(0),
		Principal:                   big.NewInt(0),
		FaultySectors:               big.NewInt(0),
		LiveSectors:                 big.NewInt(0),
		GreenScore:                  big.NewInt(0),
	}
}

func NullishVerifiableCredential(claim AgentData) (*abigen.VerifiableCredential, error) {
	claimBytes, err := AbiEncodeClaim(claim)
	if err != nil {
		return &abigen.VerifiableCredential{}, err
	}
	return &abigen.VerifiableCredential{
		Issuer:          common.Address{},
		Subject:         big.NewInt(0),
		EpochIssued:     big.NewInt(0),
		EpochValidUntil: big.NewInt(0),
		Value:           big.NewInt(0),
		Action:          [4]byte{},
		Target:          0,
		Claim:           claimBytes,
	}, nil
}

func Digest(
	chainId int64, verifyingContract common.Address, vc *abigen.VerifiableCredential,
) (common.Hash, error) {
	_, rawData, err := apitypes.TypedDataAndHash(TypedEIP712(chainId, verifyingContract, vc))
	if err != nil {
		return common.Hash{}, err
	}

	return crypto.Keccak256Hash([]byte(rawData)), nil
}
