package tui

import (
	"context"

	"github.com/AEKDA/aereq/internal/pkg/aereq"
	"github.com/AEKDA/aereq/internal/pkg/logger"
)

type tui struct {
	backend *aereq.Backend
}

func New(ctx context.Context, backend *aereq.Backend) *tui {
	return &tui{
		backend: backend,
	}
}

func (c *tui) Run(ctx context.Context) error {
	ctx = logger.WithName(ctx, "tui")

	return nil
}
