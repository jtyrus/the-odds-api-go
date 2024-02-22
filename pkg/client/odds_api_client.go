package client

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/jtyrus/the-odds-api-go/pkg/common"
	"github.com/jtyrus/the-odds-api-go/pkg/common/bookmakers"
	"github.com/jtyrus/the-odds-api-go/pkg/common/regions"
	"github.com/jtyrus/the-odds-api-go/pkg/common/sports"
	"github.com/jtyrus/the-odds-api-go/pkg/odds"
)

type HttpClient interface {
	Get(url string) (resp *http.Response, err error)
}

type OddsClient struct {
	ApiKey string
	HttpClient HttpClient
}

type OddsRequest struct {
	EventIds []string //filters for specific events
	Markets []common.Market //Default h2h
	Regions []regions.Region
	Bookmakers []bookmakers.Bookmaker
	Sport sports.Sport
	DateFormat common.DateFormat //Format of returned timestamps. Can be iso (ISO8601) or unix timestamp (seconds since epoch). Default iso
	OddsFormat common.OddsFormat //Default decimal
	TimeFrom time.Time
	TimeTo time.Time
}

const host = "https://api.the-odds-api.com/v4"

func (client *OddsClient) GetOdds(request OddsRequest) ([]odds.Event, error) {
	regionsString := ""
	for i := 0; i < len(request.Regions) - 1; i++ {
		regionsString += regions.RegionKeys[request.Regions[i]] + ","
	}
	regionsString += regions.RegionKeys[request.Regions[len(request.Regions) - 1]]
	
	url := fmt.Sprintf("%s/sports/%s/odds?regions=%s&apiKey=%s", host, request.Sport.Key, regionsString, client.ApiKey)
	url += getDates(request)
	url += getMarkets(request)
	url += getBookmakers(request)
	url += getOddsFormat(request)
	if request.EventIds != nil && len(request.EventIds) > 0 {
		url += "&eventIds=" + strings.Join(request.EventIds, ",")
	}

	client.HttpClient.Get(url)
	return nil, nil
}

func getDates(request OddsRequest) string {
	params := ""
	if request.DateFormat != "" {
		params += "&dateFormat=" + string(request.DateFormat)
	}

	if !request.TimeFrom.IsZero() {
		params += "&commenceTimeFrom=" + request.TimeFrom.Format("2006-01-02T15:04:05.000Z")
	}

	if !request.TimeTo.IsZero() {
		params += "&commenceTimeTo=" + request.TimeTo.Format("2006-01-02T15:04:05.000Z")
	}

	return params
}

func getMarkets(request OddsRequest) string {
	marketString := ""
	if len(request.Markets) != 0 {
		marketString = "&markets="
		for i := 0; i < len(request.Markets) - 1; i++ {
			marketString += string(request.Markets[i]) + ","
		}
		marketString += string(request.Markets[len(request.Markets) - 1])
	}

	return marketString
}

func getBookmakers(request OddsRequest) string {
	bookmakersString := ""
	if len(request.Bookmakers) != 0 {
		bookmakersString = "&bookmakers="
		for i := 0; i < len(request.Bookmakers) - 1; i++ {
			bookmakersString += request.Bookmakers[i].Key + ","
		}
		bookmakersString += request.Bookmakers[len(request.Bookmakers) - 1].Key
	}

	return bookmakersString
}

func getOddsFormat(request OddsRequest) string {
	formatString := ""
	if request.OddsFormat != "" {
		formatString += "&oddsFormat=" + string(request.OddsFormat)
	}

	return formatString
}