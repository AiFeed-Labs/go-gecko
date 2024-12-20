package coingecko

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"

	"github.com/AiFeed-Labs/go-gecko/format"
	"github.com/AiFeed-Labs/go-gecko/v3/types"
)

// Exchanges list, paginated
func (c *Client) Exchanges(perPage int, page int) (types.ExchangesDetail, error) {
	params := url.Values{}
	// per_page
	if perPage <= 0 || perPage > 250 {
		perPage = 100
	}
	params.Add("per_page", format.Int2String(perPage))
	params.Add("page", format.Int2String(page))

	url := fmt.Sprintf("%s/exchanges?%s", BaseURL, params.Encode())
	resp, err := c.MakeReq(url)
	if err != nil {
		return nil, err
	}
	var data types.ExchangesDetail
	err = json.Unmarshal(resp, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// Exchanges list, no pagination required
func (c *Client) ExchangesList() (types.ExchangesBase, error) {
	url := fmt.Sprintf("%s/exchanges/list", BaseURL)
	resp, err := c.MakeReq(url)
	if err != nil {
		return nil, err
	}
	var data types.ExchangesBase
	err = json.Unmarshal(resp, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// Get exchange tickers (paginated)
func (c *Client) ExchangeIDTickers(exchangeID string, coinIDs []string) (*types.ExchangeIDTickers, error) {
	params := url.Values{}
	coindIDsParam := strings.Join(coinIDs[:], ",")

	params.Add("coin_ids", coindIDsParam)

	url := fmt.Sprintf("%s/exchanges/%s/tickers?%s", BaseURL, exchangeID, params.Encode())
	resp, err := c.MakeReq(url)
	if err != nil {
		return nil, err
	}

	var data types.ExchangeIDTickers
	err = json.Unmarshal(resp, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}
