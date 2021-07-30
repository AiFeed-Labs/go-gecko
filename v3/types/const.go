package types

const (
	CoinNetworkIDBitcoin         = "bitcoin"
	CoinNetworkIDEthereum        = "ethereum"
	CoinNetworkIDEthereumClassic = "ethereum-classic"
	CoinNetworkIDBinance         = "binancecoin"
	CoinNetworkIDHuobi           = "huobi-token"
	CoinNetworkIDMatic           = "matic-network"
)

const (
	CoinDecimalsBitcoin         = 18
	CoinDecimalsEthereum        = 18
	CoinDecimalsEthereumClassic = 18
	CoinDecimalsBinance         = 18
	CoinDecimalsHuobi           = 18
	CoinDecimalsMatic           = 18
)

var CoinDecmialsMap = map[string]int{
	CoinNetworkIDBitcoin:         CoinDecimalsBitcoin,
	CoinNetworkIDEthereum:        CoinDecimalsEthereum,
	CoinNetworkIDEthereumClassic: CoinDecimalsEthereumClassic,
	CoinNetworkIDBinance:         CoinDecimalsBinance,
	CoinNetworkIDHuobi:           CoinDecimalsHuobi,
	CoinNetworkIDMatic:           CoinDecimalsMatic,
}
