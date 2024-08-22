package internal

import "github.com/codecrafters-io/grep-starter-go/cmd/mygrep/handler"

var Queue []handler.Handler = []handler.Handler{
	handler.Digit{},
	handler.Word{},
}
