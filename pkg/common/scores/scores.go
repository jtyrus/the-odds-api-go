package scores

import "github.com/jtyrus/the-odds-api-go/pkg/common"

type ScoreResponse struct {
	Scores []Score
	common.Quota
}
// List of events is returned by the /odds endpoints
type Score struct {
	Id           string      `json:"id"`
	SportKey     string      `json:"sport_key"`
	SportTitle   string      `json:"sport_title"`
	CommenceTime string      `json:"commence_time"`
	IsCompleted  bool 		 `json:"completed"`
	HomeTeam     string      `json:"home_team"`
	AwayTeam     string      `json:"away_team"`
	Scores       []TeamScores `json:"scores"`
}

type TeamScores struct {
	Name	string `json:"name"`
	Score	string `json:"score"`
}