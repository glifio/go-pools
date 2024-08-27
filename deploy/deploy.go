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

var InfinityPoolV0 = common.HexToAddress("0x43dAe5624445e7679D16a63211c5ff368681500c")
var InfinityPoolV2 = common.HexToAddress("0xe764Acf02D8B7c21d2B6A8f0a96C78541e0DC3fd")

var AgentPoliceV0 = common.HexToAddress("0x1403D2163B4beA09F1b26cA4479FFc60Ca2E46BF")
var AgentPoliceV2 = common.HexToAddress("0xE87f7d4C9f9BD235CB1b15b88Ed5E9a844726947")

var ProtoMeta = types.ProtocolMeta{
	AgentFactory:  common.HexToAddress("0x526Ab27Af261d28c2aC1fD24f63CcB3bd44D50e0"),
	AgentPolice:   AgentPoliceV2,
	IFIL:          common.HexToAddress("0x690908f7fa93afC040CFbD9fE1dDd2C2668Aa0e0"),
	InfinityPool:  InfinityPoolV2,
	MinerRegistry: common.HexToAddress("0xF246acF2b24316fE108ff906f6e90833a14e9263"),
	Router:        common.HexToAddress("0xe6e6e71E747EeD9fe45F0ff63E1B3E4a7c1199bF"),
	WFIL:          common.HexToAddress("0x60E1773636CF5E4A227d9AC24F20fEca034ee25A"),
	ChainID:       big.NewInt(314),
}

var TestInfinityPoolV0 = common.HexToAddress("0x37621D9a49FFE8f13BF2Adcb87f1b115D82ac378")
var TestInfinityPoolV2 = common.HexToAddress("0x0E5Ef6a4848d6F5eF18db8928672a29Cdc6c9a92")

var TestAgentPoliceV0 = common.HexToAddress("0xDA730B08A67D261fD8343c2061a269A2296c8469")
var TestAgentPoliceV2 = common.HexToAddress("0x55ffF1E383C02C51c06B2A041cb944fDaDe87D7c")

var TestProtoMeta = types.ProtocolMeta{
	AgentFactory:  common.HexToAddress("0x1EbF54537437a372614Eb50B6EDCbDB643603Fe7"),
	AgentPolice:   TestAgentPoliceV2,
	IFIL:          common.HexToAddress("0x8c97F94b2cDbF7Dc0098057334d9908C4dC0a885"),
	InfinityPool:  TestInfinityPoolV2,
	MinerRegistry: common.HexToAddress("0xac13FF4743594344152Af8d1626cC80257bdf5C0"),
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

var (
	TProtocolV2DeployEpoch  = big.NewInt(1912401)
	TProtocolV2UpgradeEpoch = big.NewInt(0)
)

var (
	// the agent police v2 and infinity pool v2 were deployed prior to this epoch, so it is safe to fetch events from this height
	ProtocolV2DeployEpoch = big.NewInt(4215971)
	// no events will occur on these contracts before this epoch
	ProtocolV2UpgradeEpoch = big.NewInt(0)
)
