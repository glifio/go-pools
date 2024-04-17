package util

import (
	"context"
	"crypto/ecdsa"
	"encoding/base64"
	"errors"
	"fmt"
	"log"
	"reflect"
	"regexp"
	"strings"

	"github.com/filecoin-project/go-address"
	filcrypto "github.com/filecoin-project/go-crypto"
	actorstypes "github.com/filecoin-project/go-state-types/actors"
	"github.com/filecoin-project/go-state-types/manifest"
	lotusapi "github.com/filecoin-project/lotus/api"
	"github.com/filecoin-project/lotus/chain/actors"
	"github.com/filecoin-project/lotus/chain/types"
	"github.com/filecoin-project/lotus/chain/types/ethtypes"
	"github.com/jimpick/go-ethereum/common"
	"github.com/jimpick/go-ethereum/crypto"
)

// derives a native filecoin f1 addr from a private key string
func DeriveNativeFromSECP256K1KeyString(pk string) (address.Address, error) {
	data, err := base64.StdEncoding.DecodeString(pk)
	if err != nil {
		return address.Undef, err
	}

	return address.NewSecp256k1Address(filcrypto.PublicKey(data))
}

func DeriveAddrFromPkString(pk string) (common.Address, address.Address, error) {
	pkECDSA, err := crypto.HexToECDSA(pk)
	if err != nil {
		log.Fatal(err)
	}

	return DeriveAddrFromPk(pkECDSA)
}

func DeriveAddressFromPk(pk *ecdsa.PrivateKey) (common.Address, error) {
	publicKey := pk.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return common.Address{}, fmt.Errorf("error casting public key to ECDSA")
	}

	return crypto.PubkeyToAddress(*publicKeyECDSA), nil
}

func DeriveAddrFromPk(pk *ecdsa.PrivateKey) (common.Address, address.Address, error) {
	publicKey := pk.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return common.Address{}, address.Address{}, fmt.Errorf("error casting public key to ECDSA")
	}

	evmAddr := crypto.PubkeyToAddress(*publicKeyECDSA)

	fevmAddr, err := ethtypes.ParseEthAddress(evmAddr.String())
	if err != nil {
		return common.Address{}, address.Address{}, err
	}

	delegatedAddr, err := fevmAddr.ToFilecoinAddress()
	if err != nil {
		return common.Address{}, address.Address{}, err
	}

	return evmAddr, delegatedAddr, nil
}

func DelegatedFromEthAddr(addr common.Address) (address.Address, error) {
	fevmAddr, err := ethtypes.ParseEthAddress(addr.String())
	if err != nil {
		return address.Address{}, err
	}

	return fevmAddr.ToFilecoinAddress()
}

// IsZeroAddress validate if it's a 0 address
func IsZeroAddress(address common.Address) bool {
	zeroAddressBytes := common.FromHex("0x0000000000000000000000000000000000000000")
	addressBytes := address.Bytes()
	return reflect.DeepEqual(addressBytes, zeroAddressBytes)
}

func TruncateAddr(addr string) string {
	re := regexp.MustCompile(`^[tf][0]`)
	// return ID addresses as-is
	if re.MatchString(addr) {
		return addr
	}

	firstSix := addr[:6]
	lastFour := addr[len(addr)-4:]
	return fmt.Sprintf("%s...%s", firstSix, lastFour)
}

func AddressToNative(ctx context.Context, addr string, lapi *lotusapi.FullNodeStruct, allowEVMIDs bool) (address.Address, error) {
	re := regexp.MustCompile(`^[tf][0-9]`)
	if re.MatchString(addr) {
		// user passed f0, f1, f2, f3, or f4
		filAddr, err := address.NewFromString(addr)
		if err != nil {
			return address.Undef, err
		}

		if !allowEVMIDs {
			// Note that in testing, sending to an ID actor address works ok but we still block it, as this isn't intended good behavior (passing ID addrs as representations of 0x style EVM addrs)
			if err := CheckIDNotEVMActorType(ctx, filAddr, lapi); err != nil {
				return address.Undef, err
			}
		}

		return filAddr, nil
	}

	if strings.HasPrefix(addr, "0x") {
		// user passed 0x addr or account name, convert to f4
		ethAddr, err := ethtypes.ParseEthAddress(addr)
		if err != nil {
			return address.Undef, err
		}
		return ethAddr.ToFilecoinAddress()
	}

	return address.Undef, errors.New("invalid address")
}

func AddressToEVM(ctx context.Context, addr string, lapi *lotusapi.FullNodeStruct, allowEVMIDs bool) (common.Address, error) {
	return ParseAddress(ctx, addr, lapi, allowEVMIDs)
}

// using f0 ID addresses to interact with EVM or EthAccount actor types is forbidden
func CheckIDNotEVMActorType(ctx context.Context, filAddr address.Address, lapi lotusapi.FullNode) error {
	if filAddr.Protocol() == address.ID {
		actor, err := lapi.StateGetActor(ctx, filAddr, types.EmptyTSK)
		if err != nil {
			return err
		}

		actorCodeEvm, success := actors.GetActorCodeID(actorstypes.Version(actors.LatestVersion), manifest.EvmKey)
		if !success {
			return errors.New("actor code not found")
		}
		if actor.Code.Equals(actorCodeEvm) {
			return errors.New("Cant pass an ID address of an EVM actor")
		}

		actorCodeEthAccount, success := actors.GetActorCodeID(actorstypes.Version(actors.LatestVersion), manifest.EthAccountKey)
		if !success {
			return errors.New("actor code not found")
		}
		if actor.Code.Equals(actorCodeEthAccount) {
			return errors.New("Cant pass an ID address of an Eth Account")
		}
	}

	return nil
}

func ParseAddress(ctx context.Context, addr string, lapi lotusapi.FullNode, allowEVMIDs bool) (common.Address, error) {
	if strings.HasPrefix(addr, "0x") {
		return common.HexToAddress(addr), nil
	}
	// user passed f1, f2, f3, or f4
	filAddr, err := address.NewFromString(addr)

	if err != nil {
		return common.Address{}, err
	}

	if !allowEVMIDs {
		if err := CheckIDNotEVMActorType(ctx, filAddr, lapi); err != nil {
			return common.Address{}, err
		}
	}

	if filAddr.Protocol() != address.ID && filAddr.Protocol() != address.Delegated {
		filAddr, err = lapi.StateLookupID(ctx, filAddr, types.EmptyTSK)
		if err != nil {
			return common.Address{}, err
		}
	}

	ethAddr, err := ethtypes.EthAddressFromFilecoinAddress(filAddr)
	if err != nil {
		return common.Address{}, err
	}
	return common.HexToAddress(ethAddr.String()), nil
}
