package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/truco_backend/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
