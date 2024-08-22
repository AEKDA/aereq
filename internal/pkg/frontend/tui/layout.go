package tui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func newLayout(app *tview.Application) *tview.Application {

	inputField := tview.NewInputField().
		SetLabel("Enter a number: ").
		SetFieldWidth(10).
		SetDoneFunc(func(key tcell.Key) {
			app.Stop()
		})

	return app.SetRoot(inputField, true).SetFocus(inputField)
}
