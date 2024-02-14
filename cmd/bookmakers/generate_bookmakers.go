package main

import (
	"bufio"
	"os"
	"strings"
	"text/template"

	"github.com/goccy/go-json"
	"github.com/jtyrus/the-odds-api-go/pkg/common/bookmakers"
	"github.com/jtyrus/the-odds-api-go/pkg/common/regions"
)

type Template struct {
	Varname, RegionName string
	bookmakers.Bookmaker
}

func main() {
	books := []Template{}
	addBooks(aus, &books)
	addBooks(eu, &books)
	addBooks(uk, &books)
	addBooks(us, &books)

	sortBookmakers(books)
	generate(books)
}

func generate(templateList []Template) {
	tmpl, err := template.ParseFiles("cmd/bookmakers/all_bookmakers.go.tmpl")
	check(err)

	f, err := os.Create("pkg/common/bookmakers/all_bookmakers.go")
    check(err)
	defer f.Close()
	
	w := bufio.NewWriter(f)
	defer w.Flush()

	err = tmpl.Execute(w, templateList)
	check(err)
}

	
func check(e error) {
    if e != nil {
        panic(e)
    }
}

func sortBookmakers(books []Template) {
	if len(books) <= 1 {
		return
	}
	pivot := books[len(books) - 1]
	pivotVal := strings.ToLower(pivot.Varname)
	last := 0

	for i := 0; i < len(books); i++ {
		if strings.ToLower(books[i].Varname) < pivotVal {
			books[i], books[last] = books[last], books[i]
			last++
		}
	}
	books[len(books) - 1], books[last] = books[last], books[len(books) - 1]
	sortBookmakers(books[:last])
	sortBookmakers(books[last+1:])
}

func addBooks(input string, books *[]Template) {
	var booksFromString []map[string]string
	json.Unmarshal([]byte(input), &booksFromString)
	for _, book := range booksFromString {
		varName := book["name"]
		varName = strings.ReplaceAll(varName, " ", "")
		varName = strings.Split(varName, ".")[0]
		varName = strings.Split(varName, "(")[0] + "_" + strings.ToUpper(book["region"])
		if varName[0] < 'A' {
			varName = "_" + varName
		}

		region := regions.Regions[book["region"]]
		*books = append(*books, Template{ varName, region.String(), bookmakers.Bookmaker{Key: book["key"], Name: book["name"], Region: region}})
	}
}

var aus = `
[
 {
  "region": "au",
  "key": "betfair_ex_au",
  "name": "Betfair Exchange"
 },
 {
  "region": "au",
  "key": "betr_au",
  "name": "Betr"
 },
 {
  "region": "au",
  "key": "bluebet",
  "name": "BlueBet"
 },
 {
  "region": "au",
  "key": "ladbrokes_au",
  "name": "Ladbrokes"
 },
 {
  "region": "au",
  "key": "neds",
  "name": "Neds"
 },
 {
  "region": "au",
  "key": "playup",
  "name": "PlayUp"
 },
 {
  "region": "au",
  "key": "pointsbetau",
  "name": "PointsBet (AU)"
 },
 {
  "region": "au",
  "key": "sportsbet",
  "name": "SportsBet"
 },
 {
  "region": "au",
  "key": "tab",
  "name": "TAB"
 },
 {
  "region": "au",
  "key": "topsport",
  "name": "TopSport"
 },
 {
  "region": "au",
  "key": "unibet",
  "name": "Unibet"
 }
]`

var us = `[
	{
	 "region": "us",
	 "key": "betonlineag",
	 "name": "BetOnline.ag"
	},
	{
	 "region": "us",
	 "key": "betmgm",
	 "name": "BetMGM"
	},
	{
	 "region": "us",
	 "key": "betrivers",
	 "name": "BetRivers"
	},
	{
	 "region": "us",
	 "key": "betus",
	 "name": "BetUS"
	},
	{
	 "region": "us",
	 "key": "bovada",
	 "name": "Bovada"
	},
	{
	 "region": "us",
	 "key": "draftkings",
	 "name": "DraftKings"
	},
	{
	 "region": "us",
	 "key": "fanduel",
	 "name": "FanDuel"
	},
	{
	 "region": "us",
	 "key": "lowvig",
	 "name": "LowVig.ag"
	},
	{
	 "region": "us",
	 "key": "mybookieag",
	 "name": "MyBookie.ag"
	},
	{
	 "region": "us",
	 "key": "pointsbetus",
	 "name": "PointsBet (US)"
	},
	{
	 "region": "us",
	 "key": "superbook",
	 "name": "SuperBook"
	},
	{
	 "region": "us",
	 "key": "unibet_us",
	 "name": "Unibet"
	},
	{
	 "region": "us",
	 "key": "williamhill_us",
	 "name": "William Hill (Caesars)"
	},
	{
	 "region": "us",
	 "key": "wynnbet",
	 "name": "WynnBET"
	},
	{
	 "region": "us2",
	 "key": "betparx",
	 "name": "betPARX"
	},
	{
	 "region": "us2",
	 "key": "espnbet",
	 "name": "ESPN BET"
	},
	{
	 "region": "us2",
	 "key": "fliff",
	 "name": "Fliff"
	},
	{
	 "region": "us2",
	 "key": "hardrockbet",
	 "name": "Hard Rock Bet"
	},
	{
	 "region": "us2",
	 "key": "sisportsbook",
	 "name": "SI Sportsbook"
	},
	{
	 "region": "us2",
	 "key": "tipico_us",
	 "name": "Tipico"
	},
	{
	 "region": "us2",
	 "key": "windcreek",
	 "name": "Wind Creek (Betfred PA)"
	}
   ]`

var uk = `[
	{
	 "region": "uk",
	 "key": "sport888",
	 "name": "888sport"
	},
	{
	 "region": "uk",
	 "key": "betfair_ex_uk",
	 "name": "Betfair Exchange"
	},
	{
	 "region": "uk",
	 "key": "betfair_sb_uk",
	 "name": "Betfair Sportsbook"
	},
	{
	 "region": "uk",
	 "key": "betvictor",
	 "name": "Bet Victor"
	},
	{
	 "region": "uk",
	 "key": "betway",
	 "name": "Betway"
	},
	{
	 "region": "uk",
	 "key": "boylesports",
	 "name": "BoyleSports"
	},
	{
	 "region": "uk",
	 "key": "casumo",
	 "name": "Casumo"
	},
	{
	 "region": "uk",
	 "key": "coral",
	 "name": "Coral"
	},
	{
	 "region": "uk",
	 "key": "grosvenor",
	 "name": "Grosvenor"
	},
	{
	 "region": "uk",
	 "key": "ladbrokes_uk",
	 "name": "Ladbrokes"
	},
	{
	 "region": "uk",
	 "key": "leovegas",
	 "name": "LeoVegas"
	},
	{
	 "region": "uk",
	 "key": "livescorebet",
	 "name": "LiveScore Bet"
	},
	{
	 "region": "uk",
	 "key": "matchbook",
	 "name": "Matchbook"
	},
	{
	 "region": "uk",
	 "key": "mrgreen",
	 "name": "Mr Green"
	},
	{
	 "region": "uk",
	 "key": "paddypower",
	 "name": "Paddy Power"
	},
	{
	 "region": "uk",
	 "key": "skybet",
	 "name": "Sky Bet"
	},
	{
	 "region": "uk",
	 "key": "unibet_uk",
	 "name": "Unibet"
	},
	{
	 "region": "uk",
	 "key": "virginbet",
	 "name": "Virgin Bet"
	},
	{
	 "region": "uk",
	 "key": "williamhill",
	 "name": "William Hill (UK)"
	}
   ]`
var eu = `[
	{
	 "region": "eu",
	 "key": "onexbet",
	 "name": "1xBet"
	},
	{
	 "region": "eu",
	 "key": "sport888",
	 "name": "888sport"
	},
	{
	 "region": "eu",
	 "key": "betclic",
	 "name": "Betclic"
	},
	{
	 "region": "eu",
	 "key": "betfair_ex_eu",
	 "name": "Betfair Exchange"
	},
	{
	 "region": "eu",
	 "key": "betonlineag",
	 "name": "BetOnline.ag"
	},
	{
	 "region": "eu",
	 "key": "betsson",
	 "name": "Betsson"
	},
	{
	 "region": "eu",
	 "key": "betvictor",
	 "name": "Bet Victor"
	},
	{
	 "region": "eu",
	 "key": "coolbet",
	 "name": "Coolbet"
	},
	{
	 "region": "eu",
	 "key": "everygame",
	 "name": "Everygame"
	},
	{
	 "region": "eu",
	 "key": "livescorebet_eu",
	 "name": "Livescorebet (EU)"
	},
	{
	 "region": "eu",
	 "key": "marathonbet",
	 "name": "Marathon Bet"
	},
	{
	 "region": "eu",
	 "key": "matchbook",
	 "name": "Matchbook"},
{ "region": "eu",
	 "key": "mybookieag",
	 "name": "MyBookie.ag"
	},
	{
	 "region": "eu",
	 "key": "nordicbet",
	 "name": "NordicBet"
	},
	{
	 "region": "eu",
	 "key": "pinnacle",
	 "name": "Pinnacle"
	},
	{
	 "region": "eu",
	 "key": "suprabets",
	 "name": "Suprabets"
	},
	{
	 "region": "eu",
	 "key": "unibet_eu",
	 "name": "Unibet"
	},
	{
	 "region": "eu",
	 "key": "williamhill",
	 "name": "William Hill"
	}
   ]`