package client

import (
	"io"
	"net/http"

	"github.com/goccy/go-json"
	"github.com/jtyrus/the-odds-api-go/pkg/odds"
)

type HttpClient interface {
	Get(url string) (resp *http.Response, err error)
}

type OddsClient struct {
	ApiKey string
	HttpClient HttpClient
}

const host = "https://api.the-odds-api.com/v4"

func (client *OddsClient) GetOdds(request odds.OddsRequest) ([]odds.Event, error) {
	url := host + odds.BuildOddsUrlFromRequest(request, client.ApiKey)
	
	resp, err  := client.HttpClient.Get(url)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var oddsEvents odds.OddsResponse
	err = json.Unmarshal(body, &oddsEvents)
	if err != nil {
		return nil, err
	}

	return oddsEvents.Events, nil
}

func (client *OddsClient) GetEventOdds(request odds.EventRequest) ([]odds.Event, error){
	url := host + odds.BuildsEventsUrlFromRequest(request, client.ApiKey)
	
	resp, err  := client.HttpClient.Get(url)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var oddsEvents []odds.Event
	err = json.Unmarshal(body, &oddsEvents)
	if err != nil {
		return nil, err
	}

	return oddsEvents, nil
}