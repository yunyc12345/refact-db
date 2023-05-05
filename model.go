package main

import (
	"crypto/ecdsa"
	"encoding/json"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"strings"
	"yunyc12345/refact-db/contracts"
	"yunyc12345/refact-db/utils"
)

var ChainDataMap = map[int64]*ChainData{}

type MapKey struct {
	Sequence       int64
	ReceiverAppId  int64
	MailBoxAddress string
}

func (k *MapKey) toHash() string {
	j, _ := json.Marshal(k)
	hash := common.BytesToHash(j)
	return strings.ToLower(hash.Hex())
}

type MapVal struct {
	MailerEvent       *MapValMailer
	ZkBridgeEvent     *MapValZkBridge
	IsAlreadyExecuted bool
	CreateTimeStamp   int64
}

type MapValMailer struct {
	Sequence       int64
	ReceiverAppId  int64
	MailBoxAddress string
	Sender         string
	Receipt        string
	Msg            string
	TxHash         string
}

type MapValZkBridge struct {
	Sequence       int64
	Sender         string
	SenderAppId    int64
	MailBoxAddress string
	MailerAddress  string
	TxHash         string
	Key            *Key
	Nonce          uint64
}

type Config struct {
	TestnetSkPath   string  `json:"testnet_sk_path"`
	SkPath          string  `json:"sk_path"`
	ZkbridgeAbiPath string  `json:"zkbridge_abi_path"`
	MailerAbiPath   string  `json:"mailer_abi_path"`
	Chains          []Chain `json:"chains"`
}

type Chain struct {
	Name              string `json:"name"`
	Endpoint          string `json:"endpoint"`
	ZkbridgeBridge    string `json:"zkbridge_bridge"`
	Mailer            string `json:"mailer"`
	AppChainID        int64  `json:"app_chain_id"`
	ChainID           int64  `json:"chain_id"`
	StartHeight       uint64 `json:"start_height"`
	WaitNum           int    `json:"wait_num"`
	WaitSecond        int    `json:"wait_second"`
	Testnet           bool   `json:"testnet"`
	Mailbox           string `json:"mailbox,omitempty"`
	MailerGreenfield  string `json:"mailer_greenfield"`
	MailboxGreenfield string `json:"mailbox_greenfield"`
}

type ChainData struct {
	Cli      *ethclient.Client
	Zk       *contracts.ZKBridage
	Mailer   *contracts.Mailer
	MailBox  *contracts.Mailbox
	GfMailer *contracts.Mailer
	Info     Chain
}

type Key struct {
	PrivKey *ecdsa.PrivateKey
	Address *common.Address
}

func InitGlobalKey(sks []string) (map[common.Address]*Key, error) {
	ks := make(map[common.Address]*Key)
	for _, sk := range sks {
		acc1Key, err := crypto.HexToECDSA(sk)
		address := crypto.PubkeyToAddress(acc1Key.PublicKey)
		if err != nil {
			utils.Logger.Errorln(err)
			return nil, err
		}

		k := Key{}

		k.Address = &address
		k.PrivKey = acc1Key

		ks[address] = &k
	}

	return ks, nil
}

func (c *Config) InitGlobalCliMapAndZKMap() error {
	for _, chain := range c.Chains {
		c, err := ethclient.Dial(chain.Endpoint)
		if err != nil {
			utils.Logger.Errorln(err)
			return err
		}
		zk, err := contracts.NewZKBridage(common.HexToAddress(chain.ZkbridgeBridge), c)
		if err != nil {
			utils.Logger.Errorln(err)
			return err
		}
		var mailer *contracts.Mailer
		if chain.Mailer != "" {
			mailer, err = contracts.NewMailer(common.HexToAddress(chain.Mailer), c)
			if err != nil {
				utils.Logger.Errorln(err)
				return err
			}
		}

		var gfMailer *contracts.Mailer
		if chain.MailerGreenfield != "" {
			gfMailer, err = contracts.NewMailer(common.HexToAddress(chain.MailerGreenfield), c)
			if err != nil {
				utils.Logger.Errorln(err)
				return err
			}
		}

		var mailBox *contracts.Mailbox
		if chain.Mailbox != "" {
			mailBox, err = contracts.NewMailbox(common.HexToAddress(chain.Mailbox), c)
			if err != nil {
				utils.Logger.Errorln(err)
				return err
			}
		}

		cd := ChainData{
			Cli:      c,
			Zk:       zk,
			Mailer:   mailer,
			GfMailer: gfMailer,
			Info:     chain,
			MailBox:  mailBox,
		}
		ChainDataMap[chain.AppChainID] = &cd
	}
	return nil
}
