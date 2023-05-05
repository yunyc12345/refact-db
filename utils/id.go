package utils

var (
	appIdToChainId = make(map[int64]int64)
	chainIdToAppId = make(map[int64]int64)
)

const (
	MainNetAppIdEth      = 2
	MainNetAppIdBsc      = 3
	MainNetAppIdPolygon  = 4
	MainNetAppIdAvav     = 5
	MainNetAppIdFtm      = 6
	MainNetAppIdOp       = 7
	MainNetAppIdArb      = 8
	MainNetAppIdMoonbase = 9
	MainNetAppIdEvmos    = 10
	MainNetAppIdOasis    = 11
	MainNetAppIdGnosis   = 12
	MainNetAppIdMetis    = 13

	TestNetAppIdEth      = 102
	TestNetAppIdBsc      = 103
	TestNetAppIdPolygon  = 104
	TestNetAppIdAvav     = 105
	TestNetAppIdFtm      = 106
	TestNetAppIdOp       = 107
	TestNetAppIdArb      = 108
	TestNetAppIdMoonbase = 109
	TestNetAppIdEvmos    = 110
	TestNetAppIdOasis    = 111
	TestNetAppIdGnosis   = 112
	TestNetAppIdMetis    = 113
	TestNetAppIdCombo    = 114
)

func init() {
	appIdToChainId[MainNetAppIdEth] = 1
	appIdToChainId[MainNetAppIdBsc] = 56
	appIdToChainId[MainNetAppIdPolygon] = 137
	appIdToChainId[MainNetAppIdAvav] = 43114
	appIdToChainId[MainNetAppIdFtm] = 250
	appIdToChainId[MainNetAppIdOp] = 10
	appIdToChainId[MainNetAppIdArb] = 42161
	appIdToChainId[MainNetAppIdMoonbase] = 1284
	appIdToChainId[MainNetAppIdEvmos] = 9001
	appIdToChainId[MainNetAppIdOasis] = 42262
	appIdToChainId[MainNetAppIdGnosis] = 100
	appIdToChainId[MainNetAppIdMetis] = 1088

	appIdToChainId[TestNetAppIdEth] = 5
	appIdToChainId[TestNetAppIdBsc] = 97
	appIdToChainId[TestNetAppIdPolygon] = 80001
	appIdToChainId[TestNetAppIdAvav] = 43113
	appIdToChainId[TestNetAppIdFtm] = 4002
	appIdToChainId[TestNetAppIdOp] = 420
	appIdToChainId[TestNetAppIdArb] = 421613
	appIdToChainId[TestNetAppIdMoonbase] = 1287
	appIdToChainId[TestNetAppIdEvmos] = 9000
	appIdToChainId[TestNetAppIdOasis] = 42261
	appIdToChainId[TestNetAppIdGnosis] = 300
	appIdToChainId[TestNetAppIdMetis] = 588
	appIdToChainId[TestNetAppIdCombo] = 114

	for appid, chainId := range appIdToChainId {
		chainIdToAppId[chainId] = appid
	}
}

func AppIdToChainId(appid int64) int64 {
	return appIdToChainId[appid]
}

func ChanIdToAppId(chainId int64) int64 {
	return chainIdToAppId[chainId]
}

func IsTestnetAppId(appid int64) bool {
	switch appid {
	case TestNetAppIdEth,
		TestNetAppIdBsc,
		TestNetAppIdPolygon,
		TestNetAppIdAvav,
		TestNetAppIdFtm,
		TestNetAppIdOp,
		TestNetAppIdArb,
		TestNetAppIdMoonbase,
		TestNetAppIdEvmos,
		TestNetAppIdOasis,
		TestNetAppIdGnosis,
		TestNetAppIdCombo,
		TestNetAppIdMetis:
		return true
	default:
		return false
	}
}

func IsTestChainId(chainId int64) bool {
	switch chainId {
	case appIdToChainId[TestNetAppIdEth],
		appIdToChainId[TestNetAppIdBsc],
		appIdToChainId[TestNetAppIdPolygon],
		appIdToChainId[TestNetAppIdAvav],
		appIdToChainId[TestNetAppIdFtm],
		appIdToChainId[TestNetAppIdOp],
		appIdToChainId[TestNetAppIdArb],
		appIdToChainId[TestNetAppIdMoonbase],
		appIdToChainId[TestNetAppIdEvmos],
		appIdToChainId[TestNetAppIdOasis],
		appIdToChainId[TestNetAppIdGnosis],
		appIdToChainId[TestNetAppIdCombo],
		appIdToChainId[TestNetAppIdMetis]:
		return true
	default:
		return false

	}
}

func AppIdToChainNAme(appid int64) string {
	s := ""
	switch appid {
	case MainNetAppIdEth:
		s = "eth mainnet"
	case MainNetAppIdBsc:
		s = "bsc mainnet"
	case MainNetAppIdPolygon:
		s = "polygon mainnet"
	case MainNetAppIdAvav:
		s = "avax mainnet"
	case MainNetAppIdFtm:
		s = "ftm mainnet"
	case MainNetAppIdOp:
		s = "op mainnet"
	case MainNetAppIdArb:
		s = "arbi mainnet"
	case MainNetAppIdMoonbase:
		s = "moonbase mainnet"
	case MainNetAppIdEvmos:
		s = "evmos mainnet"
	case MainNetAppIdOasis:
		s = "oasis mainnet"
	case MainNetAppIdGnosis:
		s = "gnosis mainnet"
	case MainNetAppIdMetis:
		s = "metis mainnet"

	case TestNetAppIdEth:
		s = "eth testnet"
	case TestNetAppIdBsc:
		s = "bsc testnet"
	case TestNetAppIdPolygon:
		s = "polygon testnet"
	case TestNetAppIdAvav:
		s = "avax testnet"
	case TestNetAppIdFtm:
		s = "ftm testnet"
	case TestNetAppIdOp:
		s = "op testnet"
	case TestNetAppIdArb:
		s = "arbi testnet"
	case TestNetAppIdMoonbase:
		s = "moonbase testnet"
	case TestNetAppIdEvmos:
		s = "evmos testnet"
	case TestNetAppIdOasis:
		s = "oasis testnet"
	case TestNetAppIdGnosis:
		s = "gnosis testnet"
	case TestNetAppIdMetis:
		s = "metis testnet"
	case TestNetAppIdCombo:
		s = "COMBO testnet"
	}

	return s
}
