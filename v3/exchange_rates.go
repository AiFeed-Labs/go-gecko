package coingecko

import (
	"encoding/json"
	"fmt"

	"github.com/billzong/go-gecko/v3/types"
)

// ExchangeRates https://api.coingecko.com/api/v3/exchange_rates
func (c *Client) ExchangeRates() (*types.ExchangeRatesItem, error) {
	url := fmt.Sprintf("%s/exchange_rates", baseURL)
	resp, err := c.MakeReq(url)
	if err != nil {
		return nil, err
	}
	var data *types.ExchangeRatesResponse
	err = json.Unmarshal(resp, &data)
	if err != nil {
		return nil, err
	}
	return &data.Rates, nil
}
