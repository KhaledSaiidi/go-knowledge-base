package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocationAreas(pageUrl *string) (LocationAreaResp, error) {
	endpoint := "/location-area"
	fullURL := BaseURL + endpoint
	if pageUrl != nil {
		fullURL = *pageUrl
	}

	dat, ok := c.cache.Get(fullURL)
	if ok {
		fmt.Println("Cache hit!")
		locatioAriaResp := LocationAreaResp{}
		err := json.Unmarshal(dat, &locatioAriaResp)
		if err != nil {
			return LocationAreaResp{}, err
		}
		return locatioAriaResp, nil
	}
	fmt.Println("Cache miss!")

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return LocationAreaResp{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreaResp{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return LocationAreaResp{}, fmt.Errorf("bad status code: %v", resp.StatusCode)
	}

	dat, err = io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreaResp{}, err
	}

	locatioAriaResp := LocationAreaResp{}
	err = json.Unmarshal(dat, &locatioAriaResp)
	if err != nil {
		return LocationAreaResp{}, err
	}
	c.cache.Add(fullURL, dat)

	return locatioAriaResp, nil
}
