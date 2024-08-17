package main

import (
	"context"

	"github.com/AEKDA/aereq/internal/app"
	"github.com/AEKDA/aereq/internal/pkg/aereq"
	"github.com/AEKDA/aereq/internal/pkg/frontend/tui"
)

func main() {
	ctx := context.Background()

	backend := aereq.New()

	frontend := tui.New(ctx, backend)

	app := app.New(frontend)

	app.Run(ctx)
}
