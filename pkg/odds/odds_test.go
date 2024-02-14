package odds

import (
	"testing"

	"github.com/goccy/go-json"
	"github.com/stretchr/testify/assert"
)

func Test_odds(t *testing.T) {
	var actual []Event
	err := json.Unmarshal([]byte(test), &actual)
	if err != nil {
		t.Error(err)
	}

	var expected interface{}
	err = json.Unmarshal([]byte(test), &expected)
	if err != nil {
		t.Error(err)
	}

	actualString, err := json.Marshal(actual)
	if err != nil {
		t.Error(err)
	}
	var actualInterface interface{}
	json.Unmarshal([]byte(actualString), &actualInterface)


	assert.Equal(t, expected, actualInterface)
}

var test string = `[{
	"id": "04a4ea52daa7aa07eb08e449ed7f3e92",
	"sport_key": "basketball_nba",
	"sport_title": "NBA",
	"commence_time": "2024-02-11T19:10:00Z",
	"home_team": "Miami Heat",
	"away_team": "Boston Celtics",
	"bookmakers": [
		{
			"key": "draftkings",
			"title": "DraftKings",
			"last_update": "2024-02-11T19:06:40Z",
			"markets": [
				{
					"key": "h2h",
					"last_update": "2024-02-11T19:06:40Z",
					"outcomes": [
						{
							"name": "Boston Celtics",
							"price": -310
						},
						{
							"name": "Miami Heat",
							"price": 250
						}
					]
				},
				{
					"key": "spreads",
					"last_update": "2024-02-11T19:06:40Z",
					"outcomes": [
						{
							"name": "Boston Celtics",
							"price": -110,
							"point": -8.0
						},
						{
							"name": "Miami Heat",
							"price": -110,
							"point": 8.0
						}
					]
				},
				{
					"key": "totals",
					"last_update": "2024-02-11T19:06:40Z",
					"outcomes": [
						{
							"name": "Over",
							"price": -112,
							"point": 225.0
						},
						{
							"name": "Under",
							"price": -108,
							"point": 225.0
						}
					]
				}
			]
		},
		{
			"key": "wynnbet",
			"title": "WynnBET",
			"last_update": "2024-02-11T19:06:40Z",
			"markets": [
				{
					"key": "h2h",
					"last_update": "2024-02-11T19:06:40Z",
					"outcomes": [
						{
							"name": "Boston Celtics",
							"price": -319
						},
						{
							"name": "Miami Heat",
							"price": 265
						}
					]
				},
				{
					"key": "spreads",
					"last_update": "2024-02-11T19:06:40Z",
					"outcomes": [
						{
							"name": "Boston Celtics",
							"price": -114,
							"point": -8.0
						},
						{
							"name": "Miami Heat",
							"price": -106,
							"point": 8.0
						}
					]
				},
				{
					"key": "totals",
					"last_update": "2024-02-11T19:06:40Z",
					"outcomes": [
						{
							"name": "Over",
							"price": -110,
							"point": 225.0
						},
						{
							"name": "Under",
							"price": -110,
							"point": 225.0
						}
					]
				}
			]
		}
		]}]`