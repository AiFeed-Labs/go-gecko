package coingecko

import (
	"encoding/json"
	"fmt"

	"github.com/billzong/go-gecko/v3/types"
)

// Global https://api.coingecko.com/api/v3/global
func (c *Client) Global() (*types.Global, error) {
	url := fmt.Sprintf("%s/global", BaseURL)
	resp, err := c.MakeReq(url)
	if err != nil {
		return nil, err
	}
	var data *types.GlobalResponse
	err = json.Unmarshal(resp, &data)
	if err != nil {
		return nil, err
	}
	return &data.Data, nil
}
