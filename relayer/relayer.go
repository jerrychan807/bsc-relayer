package relayer

import (
	"github.com/binance-chain/bsc-relayer/common"
	config "github.com/binance-chain/bsc-relayer/config"
	"github.com/binance-chain/bsc-relayer/executor"
	"github.com/jinzhu/gorm"
	cmn "github.com/tendermint/tendermint/libs/common"
)

type Relayer struct {
	db          *gorm.DB
	cfg         *config.Config
	bbcExecutor *executor.BBCExecutor
	bscExecutor *executor.BSCExecutor
}

func NewRelayer(db *gorm.DB, cfg *config.Config, bbcExecutor *executor.BBCExecutor, bscExecutor *executor.BSCExecutor) *Relayer {
	return &Relayer{
		db:          db,
		cfg:         cfg,
		bbcExecutor: bbcExecutor,
		bscExecutor: bscExecutor,
	}
}

func (r *Relayer) Start(startHeight uint64, curValidatorsHash cmn.HexBytes) {
	// 尝试注册成为中继者
	r.registerRelayerHub()
	// 如果competition_mode为真，bsc-relayer将监控每个区块中的跨链包裹，并尝试立即交付包裹
	if r.cfg.CrossChainConfig.CompetitionMode {
		_, err := r.cleanPreviousPackages(startHeight)
		if err != nil {
			common.Logger.Errorf("failure in cleanPreviousPackages: %s", err.Error())
		}
		go r.relayerCompetitionDaemon(startHeight, curValidatorsHash)
	} else {
		go r.relayerDaemon(curValidatorsHash)
	}
	// 更新客户端信息、主要更新区块高度
	go r.bbcExecutor.UpdateClients()
	//
	go r.bscExecutor.UpdateClients()

	go r.txTracker()
	// 自动提取奖励
	go r.autoClaimRewardDaemon()

	if len(r.cfg.BSCConfig.MonitorDataSeedList) >= 2 {
		go r.doubleSignMonitorDaemon()
	}
	go r.alert()
}
