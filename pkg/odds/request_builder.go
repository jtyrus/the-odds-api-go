package odds

type Region string

type OddsRequest struct {
	Sport   string //sport key for which to return events and odds. This is obtained from the /sports endpoint. Example : americanfootball_nfl
	ApiKey  string //Access key (40 characters). Get an API key at https://the-odds-api.com/#get-access
	Regions []string
}