package handler

import (
	"strings"

	"github.com/codecrafters-io/grep-starter-go/cmd/mygrep/app"
)

type Alternation struct{}

func (a Alternation) Matches(app *app.App) bool {
	return app.Pattern[app.Ptr.PatternL] == '('
}

func (a Alternation) Run(app *app.App) bool {
	start := app.Ptr.PatternL + 1

	for app.Pattern[app.Ptr.PatternL] != ')' {
		app.Ptr.PatternL++
	}

	raw := app.Pattern[start:app.Ptr.PatternL]
	searches := strings.Split(raw, "|")

	word := app.Line[app.Ptr.LineL : app.Ptr.LineR+1]
	for _, search := range searches {
		if string(word) == search {
			return true
		}
	}

	return false
}
