package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"
	"yunyc12345/refact-db/utils"
)

func main() {
	// 1. 加载配置文件
	// 2. 生成source fetcher
	//	2.1 获取到的数据放入map
	// 3. 生成target fetcher
	//	3.1 获取到的数据从map先获取对应的val,添加对应的`数据字段`
	// 4. 启动map处理任务,数据字段完整的就构成sql数据项
	// 5. 启动定时任务
	//	5.1 如果map中的数据超过`X分钟`且`已提交字段`为否的就被认为是未提交
	//	5.2 开始向目标链提交这些交易
	//		5.2.1 成功=>构成sql数据项
	//		5.2.2 失败:
	//			5.2.1 已提交 => 放回map,且把`已提交字段`改为真
	//			5.2.2 其他 => 其他处理

	// 1.

	utils.InitLogger(&utils.LogConf{
		Level: "info",
		Path:  "log/app.log",
	})

	cfp := os.Getenv("CONFIG_FILE_PATH")

	fmt.Printf("config path: %v\n", cfp)

	jsonFile, err := os.Open(cfp)

	if err != nil {
		panic(err)
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var config Config
	err = json.Unmarshal(byteValue, &config)
	if err != nil {
		panic(err)
	}

	err = config.InitGlobalCliMapAndZKMap()
	if err != nil {
		panic(err)
	}

	var Sks []string

	keyFile, err := os.Open(config.SkPath)
	ksByteValue, _ := ioutil.ReadAll(keyFile)

	defer jsonFile.Close()

	err = json.Unmarshal(ksByteValue, &Sks)
	if err != nil {
		panic(err)
	}

	Ks, err := InitGlobalKey(Sks)
	if err != nil {
		panic(err)
	}

	println(Ks)

	m := make(map[string]*MapVal)
	sqlList := make([]MapVal, 0)

	for _, chain := range config.Chains {
		appid := chain.AppChainID
		cd := ChainDataMap[chain.AppChainID]
		// 2.
		//if chain.Mailer != "" {
		//	mailerFetcher := NewMailerFetcher(m, appid, cd)
		//	go mailerFetcher.Run()
		//}

		if chain.MailerGreenfield != "" {
			mailerFetcher := NewMailerFetcher(m, appid, cd)
			go mailerFetcher.Run()
		}

		// 3.
		if chain.ZkbridgeBridge != "" {
			mailBoxFetcher := NewMailBoxFetcher(m, appid, cd, Ks, &sqlList)
			go mailBoxFetcher.Run()
		}
	}

	// 4.
	for {
		time.Sleep(time.Second * 2)
		//j, _ := json.Marshal(sqlList)
		//utils.Logger.Infof("sql list: %v", string(j))
		utils.Logger.Infof("sql list len: %v", len(sqlList))

		//j2, _ := json.Marshal(m)
		utils.Logger.Infof("map len: %v", len(m))
		//utils.Logger.Infof("map: %v", m)
	}

}
