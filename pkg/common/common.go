package common

type Market string

const (
	HeadToHead Market = "h2h"
	Spreads    Market = "spreads"
	Totals     Market = "totals"
	Outrights  Market = "outrights"
)

type OddsFormat string

const (
	Decimal  OddsFormat = "decimal"
	American OddsFormat = "american"
)

type DateFormat string

const (
	Iso  string = "iso"
	Unix string = "unix"
)

type Quota struct {
	Remaining int
	Used      int
}