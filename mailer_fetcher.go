package main

import (
	"context"
	"encoding/json"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"strings"
	"sync"
	"time"
	"yunyc12345/refact-db/contracts"
	"yunyc12345/refact-db/utils"
)

type MailerFetcher struct {
	wg        sync.WaitGroup
	appId     int64
	chainData *ChainData
	cancel    context.CancelFunc
	m         map[string]*MapVal
}

func NewMailerFetcher(m map[string]*MapVal, appId int64, cd *ChainData) *MailerFetcher {
	return &MailerFetcher{
		wg:        sync.WaitGroup{},
		appId:     appId,
		chainData: cd,
		cancel:    nil,
		m:         m,
	}
}

func (f *MailerFetcher) Run() {
	utils.Logger.Infof("chain name: %v, mailer fetcher run", f.chainData.Info.Name)
	ctx := context.Background()
	height := uint64(0)
	var err error
	if f.chainData.Info.StartHeight != 0 {
		height = f.chainData.Info.StartHeight
	} else {
		height, err = f.chainData.Cli.BlockNumber(ctx)
		if err != nil {
			utils.Logger.Errorf("chain: %v, get abi error :%v", f.chainData.Info.Name, err)
			return
		}
	}

	utils.Logger.Infof("monitor start: chain: %v, height: %v", f.chainData.Info.Name, height)
	mailerMetaAbi, err := contracts.MailerMetaData.GetAbi()

	if err != nil {
		utils.Logger.Errorf("chain: %v, get abi error :%v", f.chainData.Info.Name, err)
		return
	}

	event := mailerMetaAbi.Events["MessageSend"]
	for {
		//utils.Logger.Infof("chain: %v, height: %v", f.chainData.Info.Name, height)
		retry := 1
		query := ethereum.FilterQuery{
			BlockHash: nil,
			FromBlock: big.NewInt(int64(height)),
			ToBlock:   big.NewInt(int64(height)),
			Addresses: []common.Address{common.HexToAddress(f.chainData.Info.MailerGreenfield)},
			Topics:    [][]common.Hash{{event.ID}},
		}

		logs, err := f.chainData.Cli.FilterLogs(ctx, query)
		for err != nil && retry <= 10 {
			utils.Logger.Errorf("chain: %v, retry: %v, FilterLogs error :%v", f.chainData.Info.Name, retry, err)
			logs, err = f.chainData.Cli.FilterLogs(ctx, query)
			retry++
			time.Sleep(time.Duration(5) * time.Second)
		}
		if retry >= 10 && err != nil {
			utils.Logger.Errorf("chain: %v, retry: %v, err: %v", f.chainData.Info.Name, retry, err)
			panic("")
		}

		//utils.Logger.Infof("mailer log size: %v", len(logs))

		for _, l := range logs {
			seq := l.Topics[1].Big().Int64()
			receiverAppId := l.Topics[2].Big().Int64()
			mailBoxAddress := strings.ToLower(common.BytesToAddress(l.Topics[3].Bytes()).Hex())
			eventData, err := mailerMetaAbi.Unpack(event.Name, l.Data)
			if err != nil {
				utils.Logger.Errorf("chain: %v, err: %v", f.chainData.Info.Name, err)
				continue
			}

			mapK := MapKey{
				Sequence:       seq,
				ReceiverAppId:  receiverAppId,
				MailBoxAddress: mailBoxAddress,
			}
			j, _ := json.Marshal(mapK)
			utils.Logger.Infof("mailer map key: %v", string(j))
			mapV := MapVal{
				MailerEvent: &MapValMailer{
					Sequence:       seq,
					ReceiverAppId:  receiverAppId,
					MailBoxAddress: mailBoxAddress,
					Sender:         strings.ToLower(eventData[0].(common.Address).Hex()),
					Receipt:        strings.ToLower(eventData[1].(common.Address).Hex()),
					Msg:            eventData[2].(string),
					TxHash:         strings.ToLower(l.TxHash.Hex()),
				},
				ZkBridgeEvent:     nil,
				IsAlreadyExecuted: false,
				CreateTimeStamp:   time.Now().Unix(),
			}
			mapkey := mapK.toHash()
			utils.Logger.Infof("mailer mapkey: %v", mapkey)
			f.m[mapkey] = &mapV
		}

		height2, err := f.chainData.Cli.BlockNumber(ctx)
		if err != nil {
			utils.Logger.Errorf("chain: %v, get abi error :%v", f.chainData.Info.Name, err)
			return
		}
		dif := height2 - height
		if dif > 30 {
			time.Sleep(time.Duration(30/30) * time.Second)
		} else if dif <= 30 && 15 < dif {
			time.Sleep(time.Duration(30/15) * time.Second)
		} else if dif <= 15 && dif > 10 {
			time.Sleep(time.Duration(30/10) * time.Second)
		} else {
			time.Sleep(time.Duration(f.chainData.Info.WaitSecond) * time.Second)
		}

		height++
	}
}
