package deploy

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/glifio/go-pools/types"
)

var dialAddr = "https://api.node.glif.io/rpc/v1"
var t_dialAddr = "https://api.calibration.node.glif.io/rpc/v1"

var Extern = types.Extern{
	AdoAddr:       "https://ado.glif.link/rpc/v0",
	LotusDialAddr: dialAddr,
	LotusToken:    "",
}

var TestExtern = types.Extern{
	AdoAddr:       "https://ado-testnet.glif.link/rpc/v0",
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
	AgentFactory:  common.HexToAddress("0x1EbF54537437a372614Eb50B6EDCbDB643603Fe7"),
	AgentPolice:   common.HexToAddress("0xDA730B08A67D261fD8343c2061a269A2296c8469"),
	IFIL:          common.HexToAddress("0x8c97F94b2cDbF7Dc0098057334d9908C4dC0a885"),
	InfinityPool:  common.HexToAddress("0x37621D9a49FFE8f13BF2Adcb87f1b115D82ac378"),
	SimpleRamp:    common.HexToAddress("0xbf3dE3D08A8C93fFB87A02E3Cb9758746450246e"),
	MinerRegistry: common.HexToAddress("0xac13FF4743594344152Af8d1626cC80257bdf5C0"),
	PoolRegistry:  common.HexToAddress("0x9de147901aA6a87bFC6BC1c84b2fE1a60bf38e1A"),
	Router:        common.HexToAddress("0xa0770911848Bd1A0930Fc6782522381Bb6256617"),
	WFIL:          common.HexToAddress("0xaC26a4Ab9cF2A8c5DBaB6fb4351ec0F4b07356c4"),
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
