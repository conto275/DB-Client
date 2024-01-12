package stations

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

type Client struct {
	client   *http.Client
	baseUrl  string
	clientID string
	apiKey   string
}

func NewStationsClient(baseUrl string, clientID string, apiKey string, timeout time.Duration) (*Client, error) {
	if timeout == 0 {
		return nil, errors.New("timeout can't be zero ")
	}
	return &Client{
		client: &http.Client{
			Timeout: timeout,
			Transport: &loggingRoundTripper{
				logger: os.Stdout,
				next:   http.DefaultTransport,
			},
		},
		baseUrl:  baseUrl,
		clientID: clientID,
		apiKey:   apiKey,
	}, nil
}

func (c *Client) GetStation(stationID string) (Station, error) {
	endpoint := "stations"

	url := fmt.Sprintf("%s/%s/%s", c.baseUrl, endpoint, stationID)

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("DB-Client-Id", c.clientID)
	req.Header.Add("DB-Api-Key", c.apiKey)
	req.Header.Add("accept", "application/vnd.de.db.ris+json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		// Обработка ошибки
		return Station{}, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		// Обработка ошибки
		return Station{}, err
	}

	var r Station
	if err = json.Unmarshal(body, &r); err != nil {
		return Station{}, nil
	}
	return r, nil

}

// func getStations()

// func getTransportCompanies()
