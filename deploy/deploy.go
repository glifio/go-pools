package deploy

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/glifio/go-pools/types"
)

var dialAddr = "https://api.node.glif.io/rpc/v1"
var t_dialAddr = "https://api.calibration.node.glif.io/rpc/v1"

var Extern = types.Extern{
	AdoAddr:       "https://ado.glif.link/rpc/v1",
	LotusDialAddr: dialAddr,
	LotusToken:    "",
}

var TestExtern = types.Extern{
	AdoAddr:       "https://ado-testnet.glif.link/rpc/v1",
	LotusDialAddr: t_dialAddr,
	LotusToken:    "",
}

var ProtoMeta = types.ProtocolMeta{
	AgentFactory:  common.HexToAddress("0x526Ab27Af261d28c2aC1fD24f63CcB3bd44D50e0"),
	AgentPolice:   common.HexToAddress("0x1403D2163B4beA09F1b26cA4479FFc60Ca2E46BF"),
	IFIL:          common.HexToAddress("0x690908f7fa93afC040CFbD9fE1dDd2C2668Aa0e0"),
	InfinityPool:  common.HexToAddress("0x43dAe5624445e7679D16a63211c5ff368681500c"),
	SimpleRamp:    common.HexToAddress("0x8147AccE69d711bcED176a0b7a029Ff54800d930"),
	MinerRegistry: common.HexToAddress("0xF246acF2b24316fE108ff906f6e90833a14e9263"),
	PoolRegistry:  common.HexToAddress("0x7aFe4f4Ca8f301FA4177CCCDd10a4a22756588e2"),
	Router:        common.HexToAddress("0xe6e6e71E747EeD9fe45F0ff63E1B3E4a7c1199bF"),
	WFIL:          common.HexToAddress("0x60E1773636CF5E4A227d9AC24F20fEca034ee25A"),
	ChainID:       big.NewInt(314),
}

var TestProtoMeta = types.ProtocolMeta{
	AgentFactory:  common.HexToAddress("0xeb9ED18AaDd2C9e6d805328D082d4b09db9a1152"),
	AgentPolice:   common.HexToAddress("0x95e6987456a7f84A9F0355C9F501f977dA6D228d"),
	IFIL:          common.HexToAddress("0xE22C205d2d2C0b8D3215025639BBbF2984316aa1"),
	InfinityPool:  common.HexToAddress("0xF7f1905073bC968b2B0746F949e18F4c5C390Fb7"),
	SimpleRamp:    common.HexToAddress("0x0000000000000000000000000000000000000000"),
	MinerRegistry: common.HexToAddress("0xD228e88650468ac169805136ED4Df9d29dEd5F6d"),
	PoolRegistry:  common.HexToAddress("0x3a4a2A203149fccc95421ce2b966d9beA6a866ba"),
	Router:        common.HexToAddress("0xf7Aad120bF9540Cdd69Ba2ae10dC2820e8e700e3"),
	WFIL:          common.HexToAddress("0x2583b67A8dD8c252a570376c79A43095895e1d4D"),
	ChainID:       big.NewInt(314159),
}

var (
	FilForwarder  = common.HexToAddress("0x2B3ef6906429b580b7b2080de5CA893BC282c225")
	TFilForwarder = common.HexToAddress("0x2B3ef6906429b580b7b2080de5CA893BC282c225")
)

var (
	ProtocolDeployEpoch  = big.NewInt(2890957)
	TProtocolDeployEpoch = big.NewInt(559673)
)
