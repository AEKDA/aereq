package app

import (
	"context"
)

type frontend interface {
	Run(context.Context) error
}
