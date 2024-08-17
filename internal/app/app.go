package app

import (
	"context"

	"github.com/AEKDA/aereq/internal/pkg/logger"
)

type App struct {
	frontend frontend
}

func (app *App) Run(ctx context.Context) {
	ctx = logger.WithName(ctx, "app")

	if err := app.frontend.Run(ctx); err != nil {
		logger.Errorf(ctx, "can't run app: %v", err)
	}
}

func New(front frontend) *App {
	return &App{
		frontend: front,
	}
}
