package odds

import (
	"fmt"

	"github.com/jtyrus/the-odds-api-go/pkg/common"
	"github.com/jtyrus/the-odds-api-go/pkg/common/regions"
	"github.com/jtyrus/the-odds-api-go/pkg/common/sports"
)

type OddsRequest struct {
	Sport   sports.Sport //sport key for which to return events and odds. This is obtained from the /sports endpoint. Example : americanfootball_nfl
	ApiKey  string //Access key (40 characters). Get an API key at https://the-odds-api.com/#get-access
	Markets []common.Market
	Regions []regions.Region
	OddsFormat string
}

func BuildUrlFromRequest(request OddsRequest) string {
	return fmt.Sprintf("https://api.the-odds-api.com/v4/sports/%s/odds?regions=%s&markets=%s&oddsFormat=%s", request.Sport.Key, request.Regions[0], request.Markets[0], request.ApiKey)
}