package executor

import (
	relayercommon "github.com/binance-chain/bsc-relayer/common"
	"github.com/ethereum/go-ethereum/common"
)

const (
	prefixLength          = 1
	sourceChainIDLength   = 2
	destChainIDLength     = 2
	channelIDLength       = 1
	sequenceLength        = 8
	totalPackageKeyLength = prefixLength + sourceChainIDLength + destChainIDLength + channelIDLength + sequenceLength

	packageStoreName  = "ibc"
	sequenceStoreName = "sc"
	maxTryTimes       = 10

	separator                           = "::"
	CrossChainPackageEventType          = "IBCPackage"
	CorssChainPackageInfoAttributeKey   = "IBCPackageInfo"
	CorssChainPackageInfoAttributeValue = "%d" + separator + "%d" + separator + "%d" // destChainID channelID sequence

	DefaultGasPrice = 20000000000 // 20 GWei

	FallBehindThreshold          = 5
	SleepSecondForUpdateClient   = 10
	DataSeedDenyServiceThreshold = 60
)

var (
	prefixForCrossChainPackageKey = []byte{0x00}
	prefixForSequenceKey          = []byte{0xf0}

	PureHeaderSyncChannelID relayercommon.CrossChainChannelID = -1
	// 只涉及以下4个合约
	// TendermintLightClient.sol负责执行 bsc-relayer 发送过来的“同步块头”的交易，和 tmHeaderValidate 预编译合约一同工作：
	tendermintLightClientContractAddr = common.HexToAddress("0x0000000000000000000000000000000000001003")
	// 中继者奖励合约
	relayerIncentivizeContractAddr    = common.HexToAddress("0x0000000000000000000000000000000000001005")
	// RelayerHub 管理 bsc-relayer 的权限。想要运行 bsc-relayer 的人必须调用合约来存入一些 BNB 以获得授权。
	relayerHubContractAddr            = common.HexToAddress("0x0000000000000000000000000000000000001006")
	// CrossChain 跨链包预处理，通过 emit 事件发送跨链包到 BC。包预处理包括序列验证和默克尔证明验证
	crossChainContractAddr            = common.HexToAddress("0x0000000000000000000000000000000000002000")
)
