package main

import (
	"bufio"
	"io"
	"net/http"
	"os"
	"strings"
	"text/template"

	"github.com/goccy/go-json"
	"github.com/jtyrus/the-odds-api-go/pkg/common/sports"
)

type SportTemplate struct {
	Varname string
	sports.Sport
}
func main() {
	resp, err := http.Get("https://api.the-odds-api.com/v4/sports?all=true&apiKey=" + os.Args[1])
	check(err)

	if resp.StatusCode != 200 {
		panic(resp.Status)
	}

	body, err  := io.ReadAll(resp.Body)
	check(err)

	var sports []SportTemplate
	err = json.Unmarshal(body, &sports)
	check(err)

	sort(sports)
	for i := 0; i < len(sports); i++ {
		varName := strings.ReplaceAll(sports[i].Title, " ", "")
		varName = strings.ReplaceAll(varName, ".", "")
		varName = strings.ReplaceAll(varName, "-", "")
		varName = strings.ReplaceAll(varName, "'", "")
		if varName[0] < 'A' {
			varName = "_" + varName
		}
		sports[i].Varname = varName
	}

	generate(sports)
}

func sort(sports []SportTemplate) {
	if len(sports) <= 1 {
		return
	}
	end := len(sports) - 1
	pivot := strings.ToLower(sports[end].Title)
	lastMoved := 0
	for i := 0; i < len(sports); i++ {
		if strings.ToLower(sports[i].Title) < pivot {
			sports[i], sports[lastMoved] = sports[lastMoved], sports[i]
			lastMoved++
		}
	}

	sports[end], sports[lastMoved] = sports[lastMoved], sports[end]
	sort(sports[:lastMoved])
	sort(sports[lastMoved + 1:])
}

func generate(sports []SportTemplate) {
	tmpl, err := template.ParseFiles("cmd/sports/all_sports.go.tmpl")
	check(err)

	f, err := os.Create("pkg/common/sports/all_sports.go")
    check(err)
	defer f.Close()
	
	w := bufio.NewWriter(f)
	defer w.Flush()

	err = tmpl.Execute(w, sports)
	check(err)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}