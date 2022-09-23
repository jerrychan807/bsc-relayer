package executor

import (
	"encoding/hex"
	"flag"
	"fmt"

	"github.com/binance-chain/bsc-relayer/common"
	"github.com/binance-chain/go-sdk/common/types"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"testing"
)

const (
	flagConfigPath         = "config-path"
	flagConfigType         = "config-type"
	flagBBCNetworkType     = "bbc-network-type"
	flagConfigAwsRegion    = "aws-region"
	flagConfigAwsSecretKey = "aws-secret-key"
)

func initFlags() {
	flag.String(flagConfigPath, "", "config file path")
	flag.String(flagConfigType, "local_private_key", "config type, local_private_key or aws_private_key")
	flag.Int(flagBBCNetworkType, int(types.TmpTestNetwork), "Binance chain network type")
	flag.String(flagConfigAwsRegion, "", "aws region")
	flag.String(flagConfigAwsSecretKey, "", "aws secret key")

	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()
	err := viper.BindPFlags(pflag.CommandLine)
	if err != nil {
		panic(err)
	}
}

// 查询区块信息
func TestGetBlockInfo(t *testing.T) {
	initFlags()
	var height int64
	height = 50
	//fmt.Println(&height)
	//for height <= 31080 {
	//	bbcNetworkType := viper.GetInt(flagBBCNetworkType)
	//	bbcExecutor, _ := NewBBCExecutor(cfg, types.ChainNetwork(bbcNetworkType))
	//	block, _ := bbcExecutor.GetClient().Block(&height)
	//	common.Logger.Infof("block height: %d", height)
	//	common.Logger.Infof("block header ValidatorsHash: %s", block.Block.Header.ValidatorsHash.String())
	//	common.Logger.Infof("block header NextValidatorsHash: %s", block.Block.Header.NextValidatorsHash.String())
	//	height += 500
	//	time.Sleep(1 * time.Second)
	//}
	bbcNetworkType := viper.GetInt(flagBBCNetworkType)
	bbcExecutor, _ := NewBBCExecutor(cfg, types.ChainNetwork(bbcNetworkType))
	//block, _ := bbcExecutor.GetClient().Block(&height)
	//blockResults, _ := bbcExecutor.GetClient().BlockResults(&height)
	////common.Logger.Infof("block header ValidatorsHash: %s", block.Block.Header.ValidatorsHash.String())
	//
	//common.Logger.Infof("%s", blockResults)
	//fmt.Println(block.Block.Height)
	//for _, event := range blockResults.Results.EndBlock.Events {
	//	common.Logger.Infof("event: %s", event.String())
	//}

	cs, err := bbcExecutor.GetInitConsensusState(height)
	if err != nil {
		fmt.Println(err)
	}
	common.Logger.Infof("cs: %v\n", cs)
	fmt.Printf("cs: %v \n", cs)
	consensusStateBytes, err := cs.EncodeConsensusState()
	// 需要转成十六进制字符串
	newConsensusStateBytesStr := hex.EncodeToString(consensusStateBytes)
	fmt.Println("consensusStateBytesHexStr: ", newConsensusStateBytesStr)

}
