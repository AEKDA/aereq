package tui

import (
	"context"

	"github.com/AEKDA/aereq/internal/pkg/aereq"
	"github.com/AEKDA/aereq/internal/pkg/logger"
	"github.com/pkg/errors"
	"github.com/rivo/tview"
)

type tui struct {
	backend *aereq.Backend
	app     *tview.Application
}

func New(ctx context.Context, backend *aereq.Backend) *tui {
	return &tui{
		backend: backend,
		app:     tview.NewApplication(),
	}
}

func (c *tui) Run(ctx context.Context) error {
	ctx = logger.WithName(ctx, "tui")

	c.app = newLayout(c.app)

	if err := c.app.Run(); err != nil {
		return errors.Wrap(err, "tview.Run")
	}

	return nil
}
