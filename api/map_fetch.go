package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Regions struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func FetchMappedRegions(url string) (Regions, error) {
	resp, err := http.Get(url)
	if err != nil {
		return Regions{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return Regions{}, fmt.Errorf("bad status code: %v", resp.StatusCode)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return Regions{}, err
	}

	fetched := Regions{}
	err = json.Unmarshal(data, &fetched)
	if err != nil {
		return Regions{}, err
	}

	return fetched, nil
}
