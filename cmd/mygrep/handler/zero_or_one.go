package handler

import "github.com/codecrafters-io/grep-starter-go/cmd/mygrep/app"

type ZeroOrOne struct{}

func (z ZeroOrOne) Matches(app *app.App) bool {
	return app.Ptr.PatternL < app.Ptr.PatternR && app.Pattern[app.Ptr.PatternL+1] == '?'
}

func (z ZeroOrOne) Run(app *app.App) bool {
	cur := app.Pattern[app.Ptr.PatternL]

	if app.Ptr.LineL <= app.Ptr.LineR && cur == app.Line[app.Ptr.LineL] {
		app.Ptr.LineL++
	}

	app.Ptr.PatternL++

	return true
}
