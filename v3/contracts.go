package coingecko

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/billzong/go-gecko/format"
	"github.com/billzong/go-gecko/v3/types"
)

// Contract Coin info from contract address
func (c *Client) CoinContract(networkID, contract string,
	localization, tickers, marketData, communityData, developerData, sparkline bool,
) (*types.CoinsID, error) {

	if len(networkID) == 0 || len(contract) == 0 {
		return nil, fmt.Errorf("parameters are all required")
	}

	params := url.Values{}
	params.Add("localization", format.Bool2String(localization))
	params.Add("tickers", format.Bool2String(tickers))
	params.Add("market_data", format.Bool2String(marketData))
	params.Add("community_data", format.Bool2String(communityData))
	params.Add("developer_data", format.Bool2String(developerData))
	params.Add("sparkline", format.Bool2String(sparkline))
	url := fmt.Sprintf("%s/coins/%s/contract/%s?%s", BaseURL, networkID, contract, params.Encode())
	resp, err := c.MakeReq(url)
	if err != nil {
		return nil, err
	}

	var data types.CoinsID
	err = json.Unmarshal(resp, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}
