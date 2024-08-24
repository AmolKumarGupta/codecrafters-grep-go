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

	for _, search := range searches {
		if string(app.Line) == search {
			return true
		}
	}

	return false
}
