package odds

import (
	"fmt"
	"strings"
	"time"

	"github.com/jtyrus/the-odds-api-go/pkg/common"
	"github.com/jtyrus/the-odds-api-go/pkg/common/bookmakers"
	"github.com/jtyrus/the-odds-api-go/pkg/common/regions"
	"github.com/jtyrus/the-odds-api-go/pkg/common/sports"
)

type EventRequest struct {
	DateFormat common.DateFormat //Format of returned timestamps. Can be iso (ISO8601) or unix timestamp (seconds since epoch). Default iso
	EventIds []string //filters for specific events
	Sport sports.Sport
	TimeFrom time.Time
	TimeTo time.Time
}

type OddsRequest struct {
	DateFormat common.DateFormat //Format of returned timestamps. Can be iso (ISO8601) or unix timestamp (seconds since epoch). Default iso
	EventIds []string //filters for specific events
	Sport sports.Sport
	TimeFrom time.Time
	TimeTo time.Time
	Markets []common.Market //Default h2h
	Regions []regions.Region
	Bookmakers []bookmakers.Bookmaker
	OddsFormat common.OddsFormat //Default decimal
}

type ScoresRequest struct {
	DateFormat common.DateFormat //Format of returned timestamps. Can be iso (ISO8601) or unix timestamp (seconds since epoch). Default iso
	DaysFrom int //The number of days in the past from which to return completed events. Valid values are integers from 1 to 3. If this field is missing, only live and upcoming events are returned.
	EventIds []string //filters for specific events
	Sport sports.Sport
}

func BuildOddsUrlFromRequest(request OddsRequest, apiKey string) string {
	regionsString := ""
	for i := 0; i < len(request.Regions) - 1; i++ {
		regionsString += regions.RegionKeys[request.Regions[i]] + ","
	}
	regionsString += regions.RegionKeys[request.Regions[len(request.Regions) - 1]]
	
	url := fmt.Sprintf("/sports/%s/odds?regions=%s&apiKey=%s", request.Sport.Key, regionsString, apiKey)
	url += getDates(request.DateFormat, request.TimeFrom, request.TimeTo)
	url += getMarkets(request)
	url += getBookmakers(request)
	url += getOddsFormat(request)
	if request.EventIds != nil && len(request.EventIds) > 0 {
		url += "&eventIds=" + strings.Join(request.EventIds, ",")
	}

	return url
}

func BuildsEventsUrlFromRequest(request EventRequest, apiKey string) string {
	url := fmt.Sprintf("/sports/%s/events?apiKey=%s", request.Sport.Key, apiKey)
	url += getDates(request.DateFormat, request.TimeFrom, request.TimeTo)
	if request.EventIds != nil && len(request.EventIds) > 0 {
		url += "&eventIds=" + strings.Join(request.EventIds, ",")
	}

	return url
}

func BuildScoresUrlFromRequest(request ScoresRequest, apiKey string) string {
	url := fmt.Sprintf("/sports/%s/scores?apiKey=%s", request.Sport.Key, apiKey)
	url += getDates(request.DateFormat, time.UnixMicro(0), time.UnixMicro(0))
	if request.EventIds != nil && len(request.EventIds) > 0 {
		url += "&eventIds=" + strings.Join(request.EventIds, ",")
	}

	return url
}

func getDates(format common.DateFormat, from, to time.Time) string {
	params := ""
	if format != "" {
		params += "&dateFormat=" + string(format)
	}

	if !from.IsZero() {
		params += "&commenceTimeFrom=" + from.Format("2006-01-02T15:04:05.000Z")
	}

	if !to.IsZero() {
		params += "&commenceTimeTo=" + to.Format("2006-01-02T15:04:05.000Z")
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