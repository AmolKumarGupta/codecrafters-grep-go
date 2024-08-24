package handler

import "github.com/codecrafters-io/grep-starter-go/cmd/mygrep/app"

type WildCard struct{}

func (w WildCard) Matches(app *app.App) bool {
	return app.Pattern[app.Ptr.PatternL] == '.'
}

func (w WildCard) Run(app *app.App) bool {
	app.Ptr.LineL++
	return true
}
