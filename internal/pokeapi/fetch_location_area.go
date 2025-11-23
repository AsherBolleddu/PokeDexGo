package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type LocationArea struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func GetLocationArea(url string) (LocationArea, error) {
	resp, err := http.Get(url)
	if err != nil {
		return LocationArea{}, fmt.Errorf("could not get location area: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return LocationArea{}, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationArea{}, fmt.Errorf("could not read response body: %w", err)
	}

	var locationArea LocationArea
	if err := json.Unmarshal(body, &locationArea); err != nil {
		return LocationArea{}, fmt.Errorf("could not unmarshal data: %w", err)
	}

	return locationArea, nil
}
