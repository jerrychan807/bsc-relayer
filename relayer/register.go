package relayer

import (
	"time"

	"github.com/binance-chain/bsc-relayer/common"
)
// 在Bsc的RelayerHub.sol注册成为中继者
func (r *Relayer) registerRelayerHub() {
	isRelayer, err := r.bscExecutor.IsRelayer()
	if err != nil {
		panic(err)
	}
	if isRelayer {
		common.Logger.Info("This relayer has already been registered")
		return
	}

	common.Logger.Info("Register this relayer to RelayerHub")
	// 在RelayerHub合约里注册成为中继者
	_, err = r.bscExecutor.RegisterRelayer()
	if err != nil {
		panic(err)
	}
	common.Logger.Info("Waiting for register tx finalization")
	time.Sleep(20 * time.Second)

	isRelayer, err = r.bscExecutor.IsRelayer()
	if err != nil {
		panic(err)
	}
	if !isRelayer {
		panic("failed to register relayer")
	}
}
