package handler

import "github.com/codecrafters-io/grep-starter-go/cmd/mygrep/app"

type OneOrMore struct{}

func (o OneOrMore) Matches(app *app.App) bool {
	return app.Ptr.PatternL < app.Ptr.PatternR && app.Pattern[app.Ptr.PatternL+1] == '+'
}

func (o OneOrMore) Run(app *app.App) bool {
	cur := app.Pattern[app.Ptr.PatternL]

	hasOccurence := false
	for app.Ptr.LineL <= app.Ptr.LineR && cur == app.Line[app.Ptr.LineL] {
		app.Ptr.LineL++
		hasOccurence = true
	}

	app.Ptr.PatternL = app.Ptr.PatternL + 1

	return hasOccurence
}
