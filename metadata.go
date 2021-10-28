package opensea

import (
	"encoding/json"
	"net/http"
	"time"
)

type Item struct {
	Description string      `json:"description,omitempty"`
	ExternalURL string      `json:"external_url,omitempty"`
	Image       string      `json:"image,omitempty"`
	Name        string      `json:"name,omitempty"`
	Attributes  []Attribute `json:"attributes,omitempty"`
}

type Attribute struct {
	Type        string      `json:"trait_type,omitempty"`
	DisplayType string      `json:"display_type,omitempty"`
	Value       interface{} `json:"value"`
}

type API struct {
	uri string
}

func Get(tokenURI string) (*Item, error) {
	spaceClient := http.Client{
		Timeout: time.Second * 2, // Timeout after 2 seconds
	}

	req, err := http.NewRequest(http.MethodGet, tokenURI, nil)
	if err != nil {
		return nil, err
	}

	res, err := spaceClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	dec := json.NewDecoder(res.Body)

	var item Item

	err = dec.Decode(&item)

	if err != nil {
		return nil, err
	}
	return &item, nil
}
