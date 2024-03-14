package client

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/jtyrus/the-odds-api-go/pkg/common"
	"github.com/jtyrus/the-odds-api-go/pkg/common/regions"
	"github.com/jtyrus/the-odds-api-go/pkg/common/sports"
	"github.com/jtyrus/the-odds-api-go/pkg/odds"
	"github.com/stretchr/testify/assert"
)

type MockClient struct{}
func (mock *MockClient) Get(url string) (resp *http.Response, err error) {
	clientRequest = url

	return clientResponse, clientError
}
var (
	clientRequest string
	clientResponse *http.Response
	clientError error
)

func Test_GetOdds_EmptyResponse(t *testing.T) {
	oddsClient := OddsClient { HttpClient: &MockClient{}, ApiKey: "test"}
	request := odds.OddsRequest {
		Sport: sports.NBA ,
		Regions: []regions.Region { regions.UnitedStates1, regions.UnitedStates2 },
		Markets: []common.Market{ common.HeadToHead, common.Spreads, common.Totals },
		OddsFormat: common.American,
	}

	givenClientReturns("[]", 200)

	oddsClient.GetOdds(request)

	assert.Equal(t, "https://api.the-odds-api.com/v4/sports/basketball_nba/odds?regions=us,us2&apiKey=test&markets=h2h,spreads,totals&oddsFormat=american", clientRequest)
}

func Test_GetOdds_ValidResponse(t *testing.T) {
	oddsClient := OddsClient { HttpClient: &MockClient{}, ApiKey: "test"}
	request := odds.OddsRequest {
		Sport: sports.NBA,
		Regions: []regions.Region { regions.UnitedStates1, regions.UnitedStates2 },
		Markets: []common.Market{ common.HeadToHead, common.Spreads, common.Totals },
		OddsFormat: common.American,
	}

	givenClientReturns(oddsResponse(), 200)

	resp, err := oddsClient.GetOdds(request)

	assert.Nil(t, err)
	assert.Equal(t, 1, len(resp))
	assert.Equal(t, 1, len(resp[0].Bookmakers))
	assert.Equal(t, 1, len(resp[0].Bookmakers[0].Markets))
	assert.Equal(t, 2, len(resp[0].Bookmakers[0].Markets[0].Outcomes))
	assert.Equal(t, "https://api.the-odds-api.com/v4/sports/basketball_nba/odds?regions=us,us2&apiKey=test&markets=h2h,spreads,totals&oddsFormat=american", clientRequest)
}

func Test_GetEvents_SingleEventId(t *testing.T) {
	oddsClient := OddsClient { HttpClient: &MockClient{}, ApiKey: "test"}
	request := odds.EventRequest {
		Sport: sports.NBA,
		EventIds: []string { "1234" },
		DateFormat: common.DateFormat(common.Iso),
	}

	givenClientReturns(eventResponse(), 200)

	oddsClient.GetEventOdds(request)

	assert.Equal(t, "https://api.the-odds-api.com/v4/sports/basketball_nba/events?apiKey=test&dateFormat=iso&eventIds=1234", clientRequest)
}

func Test_GetEvents_MultipleEventIds(t *testing.T) {
	oddsClient := OddsClient { HttpClient: &MockClient{}, ApiKey: "test"}
	request := odds.EventRequest {
		Sport: sports.NBA,
		EventIds: []string { "1234", "0000" },
		DateFormat: common.DateFormat(common.Iso),
	}

	givenClientReturns(eventResponse(), 200)

	oddsClient.GetEventOdds(request)

	assert.Equal(t, "https://api.the-odds-api.com/v4/sports/basketball_nba/events?apiKey=test&dateFormat=iso&eventIds=1234,0000", clientRequest)
}

func givenClientReturns(body string, statusCode int) {
	clientResponse = &http.Response{
		StatusCode:    statusCode,
		Body:          io.NopCloser(bytes.NewBufferString(body)),
		ContentLength: int64(len(body)),
	}
}

func oddsResponse() string {
	return fmt.Sprint(`
	{
		"timestamp": "2023-10-10T12:10:39Z",
		"previous_timestamp": "2023-10-10T12:05:39Z",
		"next_timestamp": "2023-10-10T12:15:39Z",
		"data": [
		  {
			"id": "e912304de2b2ce35b473ce2ecd3d1502",
			"sport_key": "americanfootball_nfl",
			"sport_title": "NFL",
			"commence_time": "2023-10-11T23:10:00Z",
			"home_team": "Houston Texans",
			"away_team": "Kansas City Chiefs",
			"bookmakers": [
			  {
				"key": "draftkings",
				"title": "DraftKings",
				"last_update": "2023-10-10T12:10:29Z",
				"markets": [
				  {
					"key": "h2h",
					"last_update": "2023-10-10T12:10:29Z",
					"outcomes": [
					  {
						"name": "Houston Texans",
						"price": 2.23
					  },
					  {
						"name": "Kansas City Chiefs",
						"price": 1.45
					  }
					]
				  }
				]
			  }
			]
		  }
		]
	  }
	`);
}

func eventResponse() string {
	return `
	[
  {
    "id": "e912304de2b2ce35b473ce2ecd3d1502",
    "sport_key": "americanfootball_nfl",
    "sport_title": "NFL",
    "commence_time": "2023-10-11T23:10:00Z",
    "home_team": "Houston Texans",
    "away_team": "Kansas City Chiefs"
  }
]`
}