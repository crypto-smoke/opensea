package opensea

import (
	"encoding/json"
	"net/http"
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
	MaxValue    float64     `json:"max_value,omitempty"`
}

type API struct {
	client *http.Client
}

func New(client *http.Client) *API {
	if client == nil {
		client = http.DefaultClient
	}
	return &API{
		client: client,
	}
}

func (a *API) Get(tokenURI string) (*Item, error) {
	req, err := http.NewRequest(http.MethodGet, tokenURI, nil)
	if err != nil {
		return nil, err
	}

	res, err := a.client.Do(req)
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
