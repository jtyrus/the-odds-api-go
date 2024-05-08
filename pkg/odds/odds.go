package odds

type OddsResponse struct {
	Timestamp         string  `json:"timestamp"`
	PreviousTimestamp string  `json:"previous_timestamp"`
	NextTimestamp     string  `json:"next_timestamp"`
	Events            []Event `json:"data"`
}

// List of events is returned by the /odds endpoints
type Event struct {
	Id           string      `json:"id"`
	SportKey     string      `json:"sport_key"`
	SportTitle   string      `json:"sport_title"`
	CommenceTime string      `json:"commence_time"`
	HomeTeam     string      `json:"home_team"`
	AwayTeam     string      `json:"away_team"`
	Bookmakers   []Bookmaker `json:"bookmakers"`
}

type Bookmaker struct {
	Key        string   `json:"key"`
	Title      string   `json:"title"`
	LastUpdate string   `json:"last_update"`
	Markets    []Market `json:"markets"`
}

type Market struct {
	Key        string    `json:"key"`
	LastUpdate string    `json:"last_update"`
	Outcomes   []Outcome `json:"outcomes"`
}

type Outcome struct {
	Name  string   `json:"name"`
	Price float32  `json:"price"`
	Point *float32 `json:"point,omitempty"`
}
