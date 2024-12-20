package coingecko

import (
	"encoding/json"
	"fmt"

	"github.com/AiFeed-Labs/go-gecko/v3/types"
)

// Ping /ping endpoint
func (c *Client) Ping() (*types.Ping, error) {
	url := fmt.Sprintf("%s/ping", BaseURL)
	resp, err := c.MakeReq(url)
	if err != nil {
		return nil, err
	}
	var data *types.Ping
	err = json.Unmarshal(resp, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
