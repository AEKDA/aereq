package jsonanalyzer_test

import (
	"context"
	"testing"

	"github.com/AEKDA/aereq/internal/pkg/jsonanalyzer"
	"github.com/AEKDA/aereq/internal/pkg/logger"
)

func TestAnalyzer(t *testing.T) {
	ctx := context.Background()

	analyzer := jsonanalyzer.New()

	res, err := analyzer.Do(ctx, `{"a": 1, "b": 2}`, `.a`)
	if err != nil {
		t.Fail()
	}
	if res != "1\n" {
		logger.Errorf("res: %v", res)
		t.Fail()
	} else {
		logger.Info(res)
	}

}
