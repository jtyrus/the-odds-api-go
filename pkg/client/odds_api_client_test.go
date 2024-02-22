package client

import (
	"net/http"
	"testing"

	"github.com/jtyrus/the-odds-api-go/pkg/common"
	"github.com/jtyrus/the-odds-api-go/pkg/common/regions"
	"github.com/jtyrus/the-odds-api-go/pkg/common/sports"
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

func Test_GetOdds_Success(t *testing.T) {
	oddsClient := OddsClient { HttpClient: &MockClient{}, ApiKey: "test"}
	request := OddsRequest {
		Sport: sports.NBA,
		Regions: []regions.Region { regions.UnitedStates1, regions.UnitedStates2 },
		Markets: []common.Market{ common.HeadToHead, common.Spreads, common.Totals },
		OddsFormat: common.American,
	}
	oddsClient.GetOdds(request)

	assert.Equal(t, "https://api.the-odds-api.com/v4/sports/basketball_nba/odds?regions=us,us2&apiKey=test&markets=h2h,spreads,totals&oddsFormat=american", clientRequest)
}
