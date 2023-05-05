package main

import (
	"context"
	"encoding/json"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
	"strings"
	"sync"
	"time"
	"yunyc12345/refact-db/contracts"
	"yunyc12345/refact-db/utils"
)

type MailBoxFetcher struct {
	wg        sync.WaitGroup
	appId     int64
	chainData *ChainData
	cancel    context.CancelFunc
	m         map[string]*MapVal
	keys      map[common.Address]*Key
	sqlList   *[]MapVal
}

func NewMailBoxFetcher(m map[string]*MapVal, appId int64, cd *ChainData, keys map[common.Address]*Key, sqlList *[]MapVal) *MailBoxFetcher {
	return &MailBoxFetcher{
		wg:        sync.WaitGroup{},
		appId:     appId,
		chainData: cd,
		cancel:    nil,
		m:         m,
		keys:      keys,
		sqlList:   sqlList,
	}
}

func (f *MailBoxFetcher) Run() {
	utils.Logger.Infof("chain name: %v, mailbox fetcher run", f.chainData.Info.Name)
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
	zkbridgeAbi, err := contracts.ZKBridageMetaData.GetAbi()

	if err != nil {
		utils.Logger.Errorf("chain: %v, get abi error :%v", f.chainData.Info.Name, err)
		return
	}

	event := zkbridgeAbi.Events["ExecutedMessage"]

	for {
		//utils.Logger.Infof("chain: %v, height: %v", f.chainData.Info.Name, height)

		retry := 1
		query := ethereum.FilterQuery{
			BlockHash: nil,
			FromBlock: big.NewInt(int64(height)),
			ToBlock:   big.NewInt(int64(height)),
			Addresses: []common.Address{common.HexToAddress(f.chainData.Info.ZkbridgeBridge)},
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

		//utils.Logger.Infof("zkbridge log size: %v", len(logs))

		for _, l := range logs {
			eventData, err := zkbridgeAbi.Unpack(event.Name, l.Data)
			if err != nil {
				utils.Logger.Errorf("chain: %v, unpack log err: %v", f.chainData.Info.Name, err)
				continue
			}

			seq := l.Topics[3].Big().Int64()
			senderAppId := l.Topics[2].Big().Int64()
			mailerAddress := strings.ToLower(l.Topics[1].Hex())
			mailBoxAddress := strings.ToLower(eventData[0].(common.Address).Hex())

			mapK := MapKey{
				Sequence:       seq,
				ReceiverAppId:  f.appId,
				MailBoxAddress: mailBoxAddress,
			}
			mapKey := mapK.toHash()
			utils.Logger.Infof("zkbridge mapkey: %v", mapKey)

			tx, _, err := f.chainData.Cli.TransactionByHash(ctx, l.TxHash)
			if err != nil {
				utils.Logger.Errorf("chain: %v, get tx by hash: %v, err: %v", f.chainData.Info.Name, l.TxHash.Hex(), err)
				continue
			}

			from, err := types.Sender(types.NewEIP155Signer(tx.ChainId()), tx)
			if err != nil {
				utils.Logger.Errorf("chain: %v, parse from tx: %v, err: %v", f.chainData.Info.Name, l.TxHash.Hex(), err)
				continue
			}

			if key, ok := f.keys[from]; ok {
				if val, ok := f.m[mapKey]; ok {
					val.ZkBridgeEvent = &MapValZkBridge{
						Sequence:       seq,
						MailerAddress:  mailerAddress,
						SenderAppId:    senderAppId,
						MailBoxAddress: mailBoxAddress,
						TxHash:         strings.ToLower(l.TxHash.Hex()),
						Key:            key,
						Nonce:          tx.Nonce(),
					}

					*f.sqlList = append(*f.sqlList, *val)

					delete(f.m, mapKey)
				} else {
					j, _ := json.Marshal(mapK)
					utils.Logger.Warnf("chain: %v, map key get none: %v", f.chainData.Info.Name, string(j))
					continue
				}
			} else {
				utils.Logger.Warnf("chain: %v, from get key none: %v, tx hash: %v", f.chainData.Info.Name, from.Hex(), tx.Hash().Hex())
				continue
			}

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
