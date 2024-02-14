package regions

type Region int

const (
	UnitedKingdom Region = iota
	UnitedStates1
	UnitedStates2
	Europe
	Australia
)

var Regions = map[string]Region{
	"uk":  UnitedKingdom,
	"us":  UnitedStates1,
	"us2": UnitedStates2,
	"eu":  Europe,
	"au":  Australia,
}

var RegionKeys = map[Region]string{
	UnitedKingdom: "uk",
	UnitedStates1: "us",
	UnitedStates2: "us2",
	Europe:        "eu",
	Australia:     "au",
}