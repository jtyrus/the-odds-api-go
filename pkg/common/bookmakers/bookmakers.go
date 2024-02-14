package bookmakers

import "github.com/jtyrus/the-odds-api-go/pkg/common/regions"

type Bookmaker struct {
	Region regions.Region
	Key    string
	Name   string
}