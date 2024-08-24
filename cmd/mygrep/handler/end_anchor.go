package handler

import (
	"strings"

	"github.com/codecrafters-io/grep-starter-go/cmd/mygrep/app"
)

type EndAnchor struct{}

func (e EndAnchor) Matches(app *app.App) bool {
	return app.Pattern[app.Ptr.PatternR] == '$'
}

func (e EndAnchor) Run(app *app.App) bool {
	letter := "1234567890qwertyiuopasdfghjklzxcvbnmQWERTYIUOPASDFGHJKLZXCVBNM_"

	ptr := 1
	for app.Ptr.PatternR >= ptr && app.Ptr.LineR-1 > ptr && strings.ContainsAny(letter, string(app.Pattern[app.Ptr.PatternR-ptr])) {
		if app.Pattern[app.Ptr.PatternR-ptr] != app.Line[app.Ptr.LineR-ptr-1] {
			return false
		}

		ptr++
	}
	return true
}
