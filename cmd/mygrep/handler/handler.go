package handler

import "github.com/codecrafters-io/grep-starter-go/cmd/mygrep/app"

type Handler interface {
	Matches(matcher *app.App) bool
	Run(matcher *app.App) bool
}
