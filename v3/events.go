package coingecko

import (
	"encoding/json"
	"fmt"

	"github.com/billzong/go-gecko/v3/types"
)

// EventsCountries https://api.coingecko.com/api/v3/events/countries
func (c *Client) EventsCountries() ([]*types.EventCountryItem, error) {
	url := fmt.Sprintf("%s/events/countries", BaseURL)
	resp, err := c.MakeReq(url)
	if err != nil {
		return nil, err
	}
	var data types.EventsCountries
	err = json.Unmarshal(resp, &data)
	if err != nil {
		return nil, err
	}
	return data.Data, nil

}

// EventsTypes https://api.coingecko.com/api/v3/events/types
func (c *Client) EventsTypes() (*types.EventsTypes, error) {
	url := fmt.Sprintf("%s/events/types", BaseURL)
	resp, err := c.MakeReq(url)
	if err != nil {
		return nil, err
	}
	var data *types.EventsTypes
	err = json.Unmarshal(resp, &data)
	if err != nil {
		return nil, err
	}
	return data, nil

}
