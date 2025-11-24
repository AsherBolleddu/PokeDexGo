package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetLocationInfo(locationName string) (RespLocationInfo, error) {
	url := baseURL + "/location-area/" + locationName

	if val, ok := c.cache.Get(url); ok {
		locationInfo := RespLocationInfo{}
		err := json.Unmarshal(val, &locationInfo)
		if err != nil {
			return RespLocationInfo{}, err
		}
		return locationInfo, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespLocationInfo{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return RespLocationInfo{}, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return RespLocationInfo{}, err
	}

	locationInfo := RespLocationInfo{}
	err = json.Unmarshal(data, &locationInfo)
	if err != nil {
		return RespLocationInfo{}, err
	}

	c.cache.Add(url, data)
	return locationInfo, nil
}
