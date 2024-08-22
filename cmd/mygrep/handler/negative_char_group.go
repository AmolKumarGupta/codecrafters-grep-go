package handler

import (
	"strings"

	"github.com/codecrafters-io/grep-starter-go/cmd/mygrep/app"
)

type NegativeCharactorGroup struct{}

func (n NegativeCharactorGroup) Matches(app *app.App) bool {
	cur := app.Pattern[app.Ptr.PatternL]

	if app.Ptr.PatternL >= app.Ptr.PatternR {
		return false
	}

	nxt := app.Pattern[app.Ptr.PatternL+1]
	return cur == '[' && nxt == '^'
}

func (n NegativeCharactorGroup) Run(app *app.App) bool {
	start := app.Ptr.PatternL + 2

	for app.Ptr.PatternL <= app.Ptr.PatternR && app.Pattern[app.Ptr.PatternL] != ']' {
		app.Ptr.PatternL++
	}

	searchable := app.Pattern[start:app.Ptr.PatternL]
	line := app.Line[app.Ptr.LineL : app.Ptr.LineR+1]

	return !strings.ContainsAny(string(line), searchable)
}
